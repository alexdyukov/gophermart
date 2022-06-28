package postgres

import (
	"context"
	"database/sql"
	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
)

type GophermartStore struct {
	db *sql.DB
}

func NewGophermartStore() *GophermartStore {

	return &GophermartStore{}
}

func (p *GophermartStore) GetOrdersByUser(ctx context.Context, user string) ([]core.OrderNumber, error) {
	// work with db
	return nil, nil
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
