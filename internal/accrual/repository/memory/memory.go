package memory

import "github.com/alexdyukov/gophermart/internal/accrual/domain/core"

type AccrualStore struct {
	// map.
}

func NewAccrualStore() *AccrualStore {
	return &AccrualStore{}
}

func (m *AccrualStore) SaveOrderReceipt() error {
	// work with db.
	return nil
}

func (m *AccrualStore) SaveRewardMechanic(_ *core.Reward) error {
	// work with db.
	return nil
}

func (m *AccrualStore) GetOrderByNumber(_ int) error {
	// work with db.
	return nil
}
