package core

import (
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
	"time"
)

// UserOrderNumber is now represent users registered order.
type UserOrderNumber struct {
	Id      string
	User    string
	Status  sharedkernel.Status
	Number  int
	Accrual sharedkernel.Money
	Datе    time.Time
}

func NewOrderNumber(number int, accrual sharedkernel.Money, userID string, status sharedkernel.Status, datе time.Time) UserOrderNumber {
	return UserOrderNumber{
		Id:      sharedkernel.NewUUID(),
		User:    userID,
		Number:  number,
		Status:  status,
		Accrual: accrual,
		Datе:    datе,
	}
}
