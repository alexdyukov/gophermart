package postgres

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/alexdyukov/gophermart/internal/gophermart/auth/domain/usecase"
	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

type GophermartDB struct {
	*sql.DB
}

func NewGophermartDB(conn *sql.DB) (*GophermartDB, error) {
	dataBase := GophermartDB{
		conn,
	}

	err := dataBase.createTablesIfNotExist()
	if err != nil {
		return nil, err
	}

	return &dataBase, nil
}

func (gdb *GophermartDB) FindAllOrders(ctx context.Context, uid string) ([]core.UserOrderNumber, error) {
	result := make([]core.UserOrderNumber, 0)

	query := `SELECT uid, orderNumber, status,	accrual,dateAndTime
	FROM public.user_orders WHERE userID = $1
 	ORDER BY dateAndTime
	`

	rows, err := gdb.QueryContext(ctx, query, uid)
	// only one cuddle assignment allowed before if statement for linter
	if err != nil {
		return result, err
	}
	defer rows.Close()

	ord := core.UserOrderNumber{}

	for rows.Next() {
		err = rows.Scan(&ord.ID, &ord.Number, &ord.Status, &ord.Accrual, &ord.DateAndTime)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, usecase.ErrBadCredentials
			}

			return nil, err
		}

		result = append(result, ord)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (gdb *GophermartDB) FindAllUnprocessedOrders(ctx context.Context) ([]core.UserOrderNumber, error) {
	result := make([]core.UserOrderNumber, 0)

	query := `SELECT uid, orderNumber, userID, status, accrual, dateAndTime
	FROM public.user_orders WHERE status != $1
 	ORDER BY dateAndTime
	`
	rows, err := gdb.QueryContext(ctx, query, sharedkernel.PROCESSED)
	// only one cuddle assignment allowed before if statement for linter
	if err != nil {
		return result, err
	}
	defer rows.Close()

	ord := core.UserOrderNumber{}

	for rows.Next() {
		err = rows.Scan(&ord.ID, &ord.Number, &ord.User, &ord.Status, &ord.Accrual, &ord.DateAndTime) //#
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, sql.ErrNoRows
			}

			return nil, err
		}

		result = append(result, ord)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return result, nil
}

// nolint:funlen // ok
func (gdb *GophermartDB) FindAccountByID(ctx context.Context, userID string) (core.Account, error) {
	var ( // для сохранения чтобы потом передать в функции
		orderNumber             int64
		amount, balance         sharedkernel.Money
		operationTime           time.Time
		idWithdrawal, idAccount string
	)

	stmt, err := gdb.PrepareContext(ctx, `SELECT  uid, balance FROM user_account WHERE userID = $1`)
	if err != nil {
		return core.Account{}, err
	}

	defer func() {
		err = stmt.Close()
		if err != nil {
			log.Println("FindAccountByID :", err)
		}
	}()

	err = stmt.QueryRowContext(ctx, userID).Scan(&idAccount, &balance)
	if err != nil {
		return core.Account{}, err //nolint:wrapcheck  // ok
	}

	defer func() {
		err = stmt.Close()
		if err != nil {
			log.Println("FindAccountByID :", err)
		}
	}()

	// Withdrawals --------
	stmt, err = gdb.PrepareContext(ctx,
		`SELECT uid, orderNumber, amount, dateAndTime FROM user_withdrawals WHERE userID = $1 ORDER BY dateAndTime`)
	if err != nil {
		return core.Account{}, err
	}

	rows, err := stmt.QueryContext(ctx, userID)
	if err != nil {
		return core.Account{}, err
	}

	defer rows.Close()

	withdrawalsHistory := make([]core.AccountWithdrawals, 0)

	for rows.Next() {
		err = rows.Scan(&idWithdrawal, &orderNumber, &amount, &operationTime)
		if err != nil {
			return core.Account{}, err
		}

		accountWithdrawals := core.RestoreAccountWithdrawals(operationTime, idWithdrawal, orderNumber, amount)
		withdrawalsHistory = append(withdrawalsHistory, *accountWithdrawals)
	}

	err = rows.Err()
	if err != nil {
		return core.Account{}, err
	}

	acc := core.RestoreAccount(idAccount, userID, balance, withdrawalsHistory)

	return *acc, nil
}

// nolint:funlen // ok
func (gdb *GophermartDB) UpdateUserBalance(ctx context.Context, usrs []string) error {
	//
	var (
		userID  string
		balance sharedkernel.Money
	)

	stmt, err := gdb.PrepareContext(ctx, `SELECT SUM(accrual), userID FROM user_orders 
WHERE userID = ANY ($1) and status = $2 GROUP BY userID `)
	if err != nil {
		return err
	}

	defer func() {
		err = stmt.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	rows, err := stmt.QueryContext(ctx, usrs, sharedkernel.PROCESSED)
	if err != nil {
		return err //nolint:wrapcheck  // ok
	}

	trx, err := gdb.Begin()
	if err != nil {
		return err
	}

	defer trx.Rollback() // nolint:errcheck // ok

	for rows.Next() {
		err = rows.Scan(&balance, &userID)
		if err != nil {
			return err
		}

		err = gdb.saveToTableUserAccount(ctx, trx, sharedkernel.NewUUID(), userID, balance)
		//
		if err != nil {
			return err
		}
	}

	err = rows.Err()
	if err != nil {
		return err
	}

	err = trx.Commit() // шаг 4 — сохраняем изменения

	return err
}

func (gdb *GophermartDB) SaveUserOrder(ctx context.Context, order *core.UserOrderNumber) error {
	exists, usrID, err := orderExists(ctx, gdb.DB, order.Number)
	if err != nil || exists {
		if exists {
			if usrID != order.User {
				return sharedkernel.ErrAnotherUserOrder
			}

			return sharedkernel.ErrOrderExists
		}

		return err
	}

	trx, err := gdb.Begin()
	if err != nil {
		return err
	}

	defer trx.Rollback() // nolint:errcheck // ok

	err = gdb.saveToTableUserOrders(ctx, trx, order.ID, order.User, order.Number,
		order.Status, order.Accrual, order.DateAndTime)
	if err != nil {
		return err
	}

	err = trx.Commit()

	return err
}

func (gdb *GophermartDB) SaveOrderWithoutCheck(ctx context.Context, order *core.UserOrderNumber) error {
	trx, err := gdb.Begin()
	if err != nil {
		return err
	}

	defer trx.Rollback() // nolint:errcheck // ok

	err = gdb.saveToTableUserOrders(ctx, trx, order.ID, order.User, order.Number,
		order.Status, order.Accrual, order.DateAndTime)
	if err != nil {
		return err
	}

	err = trx.Commit()

	return err
}

func (gdb *GophermartDB) SaveAccount(ctx context.Context, acc *core.Account) error {
	balance := acc.CurrentBalance()
	uid := acc.CurrentID()
	userID := acc.CurrentUserID()

	trx, err := gdb.Begin()
	if err != nil {
		return err
	}

	defer trx.Rollback() // nolint:errcheck // ok

	err = gdb.saveToTableUserAccount(ctx, trx, uid, userID, balance)

	sliceAccountWithdrawals := core.GetSliceAccountWithdrawals(acc)

	if err != nil {
		return err
	}

	for _, withdraw := range *sliceAccountWithdrawals {
		err = gdb.saveToTableUserWithdrawals(ctx, trx, withdraw.ID, userID, withdraw.OrderNumber,
			withdraw.Amount, withdraw.OperationTime)
		if err != nil {
			return err
		}
	}

	err = trx.Commit() // шаг 4 — сохраняем изменения

	return err
}

func (gdb *GophermartDB) saveToTableUserOrders(
	ctx context.Context,
	trx *sql.Tx,
	uid string,
	userID string,
	orderNumber int64,
	status sharedkernel.Status,
	sum sharedkernel.Money,
	dateAndTime time.Time,
) error {
	//
	stmt, err := trx.PrepareContext(ctx, `
	INSERT INTO user_orders VALUES ($1, $2, $3, $4, $5, $6)
	ON CONFLICT (orderNumber, userID) DO UPDATE SET status =$4, accrual = $5, dateAndTime = $6;
	`)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, uid, orderNumber, userID,
		status, sum, dateAndTime)

	if err != nil {
		return err
	}

	return nil
}

func (gdb *GophermartDB) saveToTableUserWithdrawals(
	ctx context.Context,
	trx *sql.Tx,
	uid string,
	userID string,
	orderNumber int64,
	sum sharedkernel.Money,
	dateAndTime time.Time,
) error {
	//
	stmt, err := trx.PrepareContext(ctx, `
	INSERT INTO user_withdrawals VALUES ($1, $2, $3, $4, $5)
	ON CONFLICT (orderNumber, userID) DO NOTHING;
	`) // if we are find withdrawal don't rewrite it
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, uid, orderNumber, userID, sum, dateAndTime)
	if err != nil {
		return err
	}

	return nil
}

func (gdb *GophermartDB) saveToTableUserAccount(
	ctx context.Context,
	trx *sql.Tx,
	uid string,
	userID string,
	balance sharedkernel.Money,
) error {
	//
	stmt, err := trx.PrepareContext(ctx, `
	INSERT INTO user_account VALUES ($1, $2, $3)
	ON CONFLICT (userID) DO UPDATE SET balance =$3;`)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, uid, userID, balance)

	if err != nil {
		return err
	}

	return nil
}

func (gdb *GophermartDB) createTablesIfNotExist() error {
	_, err := gdb.Exec(`CREATE TABLE IF NOT EXISTS public.user_orders (
    											uid TEXT NOT NULL,
     											orderNumber	bigint NOT NULL, 
												userID TEXT,
												status INT  NOT NULL,
												accrual		numeric,
												dateAndTime	timestamp,
												PRIMARY KEY (userID,orderNumber)
												);
								CREATE TABLE IF NOT EXISTS public.user_withdrawals (
								    			uid TEXT NOT NULL,
     											orderNumber	bigint NOT NULL, 
												userID TEXT,
												amount		numeric,
												dateAndTime	timestamp,
												PRIMARY KEY (userID,orderNumber)
												);
	CREATE TABLE IF NOT EXISTS public.user_account (
								    			uid TEXT NOT NULL,
												userID TEXT NOT NULL,
												balance		numeric,
												PRIMARY KEY (userID)
												);
												`)
	if err != nil {
		return err // nolint:wrapcheck // ok
	}

	return nil
}

func orderExists(ctx context.Context, gdb *sql.DB, orderNumber int64) (bool, string, error) {
	//
	var (
		orderID int
		userID  string
	)

	const selectSQL = `
SELECT orderNumber, userID FROM public.user_orders WHERE orderNumber = $1;`

	err := gdb.QueryRowContext(ctx, selectSQL, orderNumber).Scan(&orderID, &userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, userID, nil
		}

		return false, userID, err
	}

	return true, userID, nil
}
