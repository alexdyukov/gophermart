package postgres

import (
	"database/sql"

	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
)

type GophermartStore struct {
	db *sql.DB
}

func NewGophermartStore() *GophermartStore {
	return &GophermartStore{
		// set up db
	}
}

func (p *GophermartStore) GetOrdersByUser(user string) []core.OrderNumber {
	// work with db
	return nil
}

func (p *GophermartStore) GetAccountByID(id string) (core.Account, error) {
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
