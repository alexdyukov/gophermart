package postgres

import (
	"database/sql"
	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
)

type GophermartStore struct {
	db *sql.DB
}

func NewDB(dataSourceName string) (*GophermartStore, error) {

	dataBase, err := sql.Open("pgx", dataSourceName)
	if err != nil {
		return nil, err
	}

	db := GophermartStore{db: dataBase}
	db.initSchema()

	return &db, nil
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

func (p *GophermartStore) initSchema() error {
	_, err := p.db.Exec(photosSchema)
	return err
}
