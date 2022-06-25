package memory

import "github.com/alexdyukov/gophermart/internal/accrual/domain/core"

type AccrualStore struct {
	// map
}

func NewAccrualStore() *AccrualStore {
	return &AccrualStore{}
}

func (m *AccrualStore) SaveCalculationProcessData() error {
	// work with db
	return nil
}

func (m *AccrualStore) SaveMechanic(mechanic core.RewardMechanic) error {
	// work with db
	return nil
}

func (m *AccrualStore) GetLoyaltyPointsByOrderNumber(number int) error {
	// work with db
	return nil
}
