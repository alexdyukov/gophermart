package memory

import (
	"context"
	"github.com/alexdyukov/gophermart/internal/accrual/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

type AccrualStore struct {
	orders map[int]core.Answer // тут индекс это номер заказа, мы его храним тут для того чтобы легче было получать.
	// Но для отправки ответа будем его еще и в структуре хранить
}

func NewAccrualStore() *AccrualStore {
	// заполним тестовыми данными чтобы можно было проверить, потом надо будет удалить.
	m := make(map[int]core.Answer)

	number := 2377225624
	av := core.Answer{Number: number, Status: sharedkernel.PROCESSED.String(), Accrual: 500}
	m[number] = av

	return &AccrualStore{orders: m}
}

func (m *AccrualStore) SaveCalculationProcessData() error {

	// work with db
	return nil
}

func (m *AccrualStore) SaveMechanic(mechanic core.RewardMechanic) error {
	// work with db
	return nil
}

func (m *AccrualStore) GetLoyaltyPointsByOrderNumber(ctx context.Context, number int) (core.Answer, error) {

	if val, inMap := m.orders[number]; inMap {
		return val, nil
	}

	return core.Answer{Number: number, Status: sharedkernel.INVALID.String()}, nil
}
