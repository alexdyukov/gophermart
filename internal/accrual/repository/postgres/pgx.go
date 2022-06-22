package postgres

import (
	"database/sql"

	"github.com/alexdyukov/gophermart/internal/accrual/domain/core"
)

type PgxDBAccrual struct {
	db *sql.DB
}

func (p *PgxDBAccrual) SaveCalculationProcessData() {
	// work with db
}

func (p *PgxDBAccrual) SaveMechanic(mechanic core.RewardMechanic) error {
	// work with db
	return nil
}

func (p *PgxDBAccrual) GetLoyaltyPointsByOrderNumber(number int) error {
	// work with db
	return nil
}
