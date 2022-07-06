package postgres

import (
	"context"
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

func (db *AccrualDB) SaveOrderReceipt(ctx context.Context, _ *core.OrderReceipt) error {
	// work with db
	return nil
}

func (db *AccrualDB) SaveRewardMechanic(ctx context.Context, _ *core.Reward) error {
	// work with db
	return nil
}

func (db *AccrualDB) GetOrderByNumber(ctx context.Context, _ int) (core.OrderReceipt, error) {
	order := core.OrderReceipt{ // fake
		Status:      sharedkernel.PROCESSING,
		Accrual:     16,     // nolint:gomnd // temporary fake
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
