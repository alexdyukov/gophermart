package memory

import (
	"context"
	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
	"time"
)

// Store ...
type GophermartStore struct {
	users       map[int]UserModel
	orders      map[int]OrderModel
	withdrawals map[int]WithdrawalModel
}

func NewGophermartStore() *GophermartStore {

	return &GophermartStore{
		users:       map[int]UserModel{},
		orders:      map[int]OrderModel{},
		withdrawals: map[int]WithdrawalModel{},
	}
}

func (p *GophermartStore) GetOrdersByUser(_ context.Context, user string) ([]core.OrderNumber, error) {

	mTest := make(map[int]OrderModel)

	av1 := OrderModel{Number: "number1", UserID: "1", Status: sharedkernel.PROCESSED, Sum: 500, Date: time.Now()}
	mTest[1] = av1
	av2 := OrderModel{Number: "number2", UserID: "1", Status: sharedkernel.NEW, Date: time.Now()}
	av3 := OrderModel{Number: "number3", UserID: "1", Status: sharedkernel.PROCESSED, Sum: 300, Date: time.Now()}
	mTest[2] = av2
	mTest[3] = av3

	p.orders = mTest

	rez := make([]core.OrderNumber, 0)
	for _, ord := range p.orders {
		if ord.UserID == user {

			a := core.OrderNumber{Id: sharedkernel.NewUUID(), User: user, Number: ord.Number, Status: ord.Status, Accrual: ord.Sum}
			rez = append(rez, a)
		}
	}

	return rez, nil
}

func (p *GophermartStore) GetAccountByID(_ context.Context, id string) (core.Account, error) {
	// по имени пользователя получаем сумму накопленных балов

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
