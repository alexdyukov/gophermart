package memory

import "github.com/alexdyukov/gophermart/internal/accrual/domain/core"

type AccrualMemoryStore struct {
	// map
}

func NewAccrualMemoryStore() *AccrualMemoryStore {
	return &AccrualMemoryStore{}
}

func (m *AccrualMemoryStore) SaveCalculationProcessData() error {
	// work with db
	return nil
}

func (m *AccrualMemoryStore) SaveMechanic(mechanic core.RewardMechanic) error {
	// work with db
	return nil
}

func (m *AccrualMemoryStore) GetLoyaltyPointsByOrderNumber(number int) error {
	// work with db
	return nil
}
