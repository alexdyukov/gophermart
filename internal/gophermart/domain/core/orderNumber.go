package core

import (
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
	"time"
)

// UserOrderNumber is now represent users registered order.
type UserOrderNumber struct {
	Id          string
	User        string
	Status      sharedkernel.Status
	Number      int
	Accrual     float32
	DateAndTime time.Time
}

func NewOrderNumber(number int, accrual float32, userID string, status sharedkernel.Status) UserOrderNumber {
	return UserOrderNumber{
		Id:      sharedkernel.NewUUID(),
		User:    userID,
		Number:  number,
		Status:  status,
		Accrual: accrual,
	}
}
