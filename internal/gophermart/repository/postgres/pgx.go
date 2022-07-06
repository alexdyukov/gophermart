package postgres

import (
	"context"
	"database/sql"

	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
)

type GophermartDB struct {
	*sql.DB
}

func NewGophermartDB(conn *sql.DB) *GophermartDB {
	return &GophermartDB{ // nolint:exhaustivestruct // ok.
		conn,
	}
}


func (p *GophermartDB) FindAllOrders(ctx context.Context, _ string) ([]core.UserOrderNumber, error) {
	// retrieve from database all user's order numbers with batched query
	// and construct list of entities
	return nil, nil
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

// func (p *GophermartDB) createTableIFNotExists() {
//	// create table for [core.UserOrderNumber] entities
//}
