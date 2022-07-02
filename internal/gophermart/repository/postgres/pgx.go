package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
	"time"
)

type GophermartStore struct {
	db *sql.DB
}

func NewGophermartStore() *GophermartStore {

	return &GophermartStore{}
}

func (p *GophermartStore) GetOrdersByUser(ctx context.Context, user string) ([]core.OrderNumber, error) {

	var (
		orderNumber string
		status      int
		accrual     float32
		dateAndTime time.Time
	)

	selectSQL := `
	SELECT 
	 orderNumber,
	 status,
	 accrual,
	 dateAndTime
	FROM orders 
	WHERE userID=$1 ;
	`
	answer := make([]core.OrderNumber, 0)

	if p.db == nil {
		return answer, errors.New("can't open db") // пустой список байт
	}
	check := new(string)
	stmt, err := p.db.PrepareContext(ctx, selectSQL)
	if err != nil {
		return answer, err
	}
	row := stmt.QueryRow(user)
	rez := make([]core.OrderNumber, 0)

	if err := row.Scan(check); err != sql.ErrNoRows {

		p.db.QueryRow(selectSQL, user).Scan(&orderNumber, &status, &accrual, &dateAndTime)

		a := core.OrderNumber{Id: sharedkernel.NewUUID(), User: user, Number: orderNumber, Status: sharedkernel.Status(status), Accrual: accrual, Data: dateAndTime}
		rez = append(rez, a)

		return rez, nil

	}

	return rez, err
}

func (p *GophermartStore) GetAccountByID(ctx context.Context, id string) (core.Account, error) {
	// work with db
	return core.Account{}, nil
}

func (p *GophermartStore) SaveOrderNumber(core.OrderNumber) error {
	// work with db
	return nil
}

func (p *GophermartStore) SaveAccount(core.Account) error {
	// work with db
	return nil
}
