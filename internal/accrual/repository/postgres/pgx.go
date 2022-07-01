package postgres

import (
	"database/sql"

	"github.com/alexdyukov/gophermart/internal/accrual/domain/core"
)

type PgxDBAccrual struct {
	*sql.DB
}

func (p *PgxDBAccrual) SaveCalculationProcessData() {
	// work with db
}

func (p *PgxDBAccrual) SaveMechanic(_ *core.RewardMechanic) error {
	// work with db
	return nil
}

func (p *PgxDBAccrual) GetLoyaltyPointsByOrderNumber(_ int) error {
	// work with db
	return nil
}
