package postgres

import (
	"database/sql"

	"github.com/alexdyukov/gophermart/internal/accrual/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

type AccrualDB struct {
	*sql.DB
}

func NewAccrualDB(conn *sql.DB) *AccrualDB {
	return &AccrualDB{
		conn,
	}
}

func (db *AccrualDB) SaveOrderReceipt(orderReceipt *core.OrderReceipt) error {
	// work with db
	return nil
}

func (db *AccrualDB) SaveRewardMechanic(_ *core.Reward) error {
	// work with db
	return nil
}

func (db *AccrualDB) GetOrderByNumber(_ int) (core.OrderReceipt, error) {
	// query data from database using income number
	// construct aggregate

	// fake
	order := core.OrderReceipt{
		Status:      sharedkernel.PROCESSING,
		Accrual:     0,
		OrderNumber: 122937, // nolint:gomnd // temporary fake
		Goods: []core.Product{
			{
				Description: "TV",
				Price:       90, // nolint:gomnd // temporary fake
			},
		},
	}

	return order, nil
}
