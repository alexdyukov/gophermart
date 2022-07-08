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

	query := `
	SELECT
	uid,
	orderNumber,
	status,
	accrual,
	dateAndTime
	FROM user_orders
	WHERE userID = $1
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

	return result, nil
}

// nolint:funlen // ok
func (gdb *GophermartDB) FindAccountByID(ctx context.Context, userID string) (core.Account, error) {
	// retrieve User's account from database and construct it with core.RestoreAccount
	var ( // для сохранения чтобы потом передать в функции
		order           int
		amount, accrual sharedkernel.Money
		operationTime   time.Time
	)
	//stmt, err := gdb.PrepareContext(ctx, `SELECT SUM(accrual) FROM user_orders WHERE userID = $1`)
	stmt, err := gdb.PrepareContext(ctx, `SELECT SUM(accrual) FROM user_orders WHERE userID = $1 and status = $2`)
	if err != nil {
		return core.Account{}, err
	}

	defer func() {
		err = stmt.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	err = stmt.QueryRowContext(ctx, userID, sharedkernel.PROCESSED).Scan(&accrual)
	//err = stmt.QueryRowContext(ctx, userID).Scan(&accrual)
	if err != nil {
		return core.Account{}, err //nolint:wrapcheck  // ok
	}

	defer func() {
		err = stmt.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	acc := core.RestoreAccount(sharedkernel.NewUUID(), userID, accrual)

	// Withdrawals --------
	stmt, err = gdb.PrepareContext(ctx,
		`SELECT orderNumber, amount, dateAndTime FROM user_withdrawals WHERE userID = $1 ORDER BY dateAndTime`)
	if err != nil {
		return *acc, err
	}

	rows, err := stmt.QueryContext(ctx, userID)
	if err != nil {
		return *acc, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&order, &amount, &operationTime)
		if err != nil {
			return *acc, err
		}

		err = acc.WithdrawPoints(order, amount, operationTime)
		if err != nil {
			return *acc, err
		}
	}

	return *acc, nil
}

func (gdb *GophermartDB) SaveUserOrder(context.Context, core.UserOrderNumber) error {
	// we receive newly created user order, and save in into db
	// return err if something goes wrong
	return nil
}

func (gdb *GophermartDB) SaveAccount(context.Context, core.Account) error {
	// Store core.Account into database
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
												PRIMARY KEY (uid)
												);

								CREATE TABLE IF NOT EXISTS public.user_withdrawals (
								    			uid TEXT NOT NULL,
     											orderNumber	bigint NOT NULL, 
												userID TEXT,
												amount		numeric,
												dateAndTime	timestamp,
												PRIMARY KEY (uid)
												);
												`)
	if err != nil {
		return err // nolint:wrapcheck // ok
	}

	return nil
}
