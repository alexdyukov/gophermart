package core

import "github.com/alexdyukov/gophermart/internal/sharedkernel"

// Move this to accrual ??

type OrderNumber struct {
	id      string
	user    string // owner
	number  string
	status  sharedkernel.Status
	accrual int
}

func NewOrderNumber(number int) OrderNumber {
	return OrderNumber{
		status: sharedkernel.NEW,
		// ...
	}
}
