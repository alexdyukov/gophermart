package core

import (
	"time"

	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

// UserOrderNumber is now represent users registered order.
type UserOrderNumber struct {
	DateAndTime time.Time
	ID          string
	User        string
	Status      sharedkernel.Status
	Number      int
	Accrual     sharedkernel.Money
}

func NewOrderNumber(number int, accrual sharedkernel.Money, userID string, status sharedkernel.Status, datе time.Time) UserOrderNumber {
	return UserOrderNumber{
		ID:          sharedkernel.NewUUID(),
		User:        userID,
		Number:      number,
		Status:      status,
		Accrual:     accrual,
		DateAndTime: datе,
	}
}
