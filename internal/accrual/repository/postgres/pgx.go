package postgres

import (
	"database/sql"

	"github.com/alexdyukov/gophermart/internal/accrual/domain/core"
)

type AccrualDB struct {
	*sql.DB
}

func NewAccrualDB(conn *sql.DB) *AccrualDB {
	return &AccrualDB{
		conn,
	}
}

func (p *AccrualDB) SavePurchasedOrder() error {
	// work with db
	return nil
}

func (p *AccrualDB) SaveMechanic(_ *core.RewardMechanic) error {
	// work with db
	return nil
}

func (p *AccrualDB) GetLoyaltyPointsByOrderNumber(_ int) error {
	// work with db
	return nil
}
