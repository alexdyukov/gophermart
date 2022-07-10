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
	Number      int64
	Accrual     sharedkernel.Money
}

func NewOrderNumber(num int64, sum sharedkernel.Money, uID string, sts sharedkernel.Status) UserOrderNumber {
	return UserOrderNumber{
		DateAndTime: time.Now(),
		ID:          sharedkernel.NewUUID(),
		User:        uID,
		Number:      num,
		Status:      sts,
		Accrual:     sum,
	}
}
