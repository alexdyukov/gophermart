package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/alexdyukov/gophermart/internal/sharedkernel"

	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
)

type GophermartDB struct {
	*sql.DB
}

func NewGophermartDB(conn *sql.DB) (*GophermartDB, error) {
	dataBase := GophermartDB{
		conn,
	}

	err := dataBase.createOrdersTableIfNotExist()
	if err != nil {
		return nil, err
	}

	return &dataBase, nil
}

func (p *GophermartDB) FindAllOrders(ctx context.Context, uid string) ([]core.UserOrderNumber, error) {
	// retrieve from database all user's order numbers with batched query
	// and construct list of entities

	rez := make([]core.UserOrderNumber, 0)

	// ord := core.NewOrderNumber(3283027263, 500.79, "", sharedkernel.NEW, time.Now())
	// rez = append(rez, ord)
	//
	// ord = core.NewOrderNumber(3283027263, 500.79, "", sharedkernel.NEW, time.Now())
	// rez = append(rez, ord)
	//
	// return rez, nil

	selectSQL := `
	SELECT
	orderNumber,
	status,
	accrual,
	dateAndTime
	FROM orders
	WHERE uid = $1
	`

	rows, err := p.QueryContext(ctx, selectSQL, uid)
	if err != nil {
		return rez, err
	}

	defer rows.Close()

	var (
		number      int
		status      int
		accrual     sharedkernel.Money
		dateAndTime time.Time
	)
	for rows.Next() {
		err = rows.Scan(&number, &status, &accrual, &dateAndTime)
		if err != nil {
			return nil, err
		}

		ord := core.UserOrderNumber{
			ID:          sharedkernel.NewUUID(),
			User:        uid,
			Number:      number,
			Status:      sharedkernel.Status(status),
			Accrual:     accrual,
			DateAndTime: dateAndTime,
		}

		rez = append(rez, ord)
	}

	return rez, nil
}

func (p *GophermartDB) FindAccountByID(ctx context.Context, _ string) (core.Account, error) {
	// retrieve User's account from database and construct it with core.RestoreAccount
	return core.Account{}, nil
}

func (p *GophermartDB) SaveUserOrder(context.Context, core.UserOrderNumber) error {
	// we receive newly created user order, and save in into db
	// return err if something goes wrong
	return nil
}

func (p *GophermartDB) SaveAccount(context.Context, core.Account) error {
	// Store core.Account into database
	return nil
}

func (p *GophermartDB) createOrdersTableIfNotExist() error {
	_, err := p.Exec(`CREATE TABLE IF NOT EXISTS public.orders (
     											orderNumber	INT NOT NULL, 
												uid TEXT,
												status INT  NOT NULL,
												accrual		numeric,
												dateAndTime	timestamp,
												PRIMARY KEY (orderNumber, uid)
												);
												`)
	if err != nil {
		return err // nolint:wrapcheck // ok
	}

	return nil
}

func (p *GophermartDB) createWithdrawalsTableIfNotExist() error {
	_, err := p.Exec(`CREATE TABLE IF NOT EXISTS public.withdrawals (
     											orderNumber	INT NOT NULL, 
												uid TEXT,
												amount		numeric,
												dateAndTime	timestamp,
												PRIMARY KEY (orderNumber, uid)
												);
												`)
	if err != nil {
		return err // nolint:wrapcheck // ok
	}

	return nil
}
