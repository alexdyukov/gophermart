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

	av1 := OrderModel{Number: "number1", UserID: "1", Status: sharedkernel.PROCESSED, Sum: 500, Date: time.Date(2022, time.May, 15, 17, 45, 12, 0, time.Local)}
	mTest[1] = av1
	av2 := OrderModel{Number: "number2", UserID: "1", Status: sharedkernel.NEW, Date: time.Date(2021, time.May, 15, 17, 45, 12, 0, time.Local)}
	av3 := OrderModel{Number: "number3", UserID: "1", Status: sharedkernel.PROCESSED, Sum: 300, Date: time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)}
	mTest[2] = av2
	mTest[3] = av3

	p.orders = mTest

	rez := make([]core.OrderNumber, 0)
	for _, ord := range p.orders {
		if ord.UserID == user {

			a := core.OrderNumber{Id: sharedkernel.NewUUID(), User: user, Number: ord.Number, Status: ord.Status, Accrual: ord.Sum, Data: ord.Date}
			rez = append(rez, a)
		}
	}

	return rez, nil
}

func (p *GophermartStore) GetAccountByID(_ context.Context, id string) (core.Account, error) {

	return core.Account{}, nil
}

func (p *GophermartStore) GetBalance(_ context.Context, user string) (float32, float32, error) {

	mOrderTest := make(map[int]OrderModel)

	av1 := OrderModel{Number: "number1", UserID: "1", Status: sharedkernel.PROCESSED, Sum: 500, Date: time.Now()}
	mOrderTest[1] = av1
	av2 := OrderModel{Number: "number2", UserID: "1", Status: sharedkernel.NEW, Date: time.Now()}
	av3 := OrderModel{Number: "number3", UserID: "1", Status: sharedkernel.PROCESSED, Sum: 300, Date: time.Now()}
	mOrderTest[2] = av2
	mOrderTest[3] = av3

	mWithdrawalrTest := make(map[int]WithdrawalModel)

	aw1 := WithdrawalModel{Number: "number4", UserID: "1", Sum: 300, Date: time.Now()}
	mWithdrawalrTest[1] = aw1
	p.orders = mOrderTest
	p.withdrawals = mWithdrawalrTest

	var sumCurr float32
	var sumWithdwr float32

	for _, ord := range p.orders {
		if ord.UserID == user {
			sumCurr += ord.Sum
		}
	}

	for _, with := range p.withdrawals {
		if with.UserID == user {
			sumWithdwr += with.Sum
		}
	}

	// work with db
	return sumCurr, sumWithdwr, nil
}

func (p *GophermartStore) SaveOrderNumber(core.OrderNumber) error {
	// work with db
	return nil
}

func (p *GophermartStore) SaveAccount(core.Account) error {
	// work with db
	return nil
}
