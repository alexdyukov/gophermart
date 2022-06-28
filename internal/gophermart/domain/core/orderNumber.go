package core

import (
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
	"time"
)

// Move this to accrual ??

type OrderNumber struct {
	Id      string
	User    string // owner
	Number  string
	Status  sharedkernel.Status
	Accrual int
	Data    time.Time
}

func NewOrderNumber(number int) OrderNumber {
	return OrderNumber{
		Status: sharedkernel.NEW,
		// ...
	}
}
