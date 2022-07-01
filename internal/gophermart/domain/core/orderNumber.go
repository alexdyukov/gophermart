package core

import "github.com/alexdyukov/gophermart/internal/sharedkernel"

// OrderNumber is now represent more Accrual order number than Gophermart
// we still did not get an answer from mentor, so functionality is not totally clear
// Fields commented because of linter checks.
type OrderNumber struct {
	id string
	// user string
	status sharedkernel.Status
	number int
	// accrual int
}

func NewOrderNumber(number int) OrderNumber {
	return OrderNumber{
		id:     sharedkernel.NewUUID(),
		number: number,
		status: sharedkernel.NEW,
	}
}
