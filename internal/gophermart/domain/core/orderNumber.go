package core

import "github.com/alexdyukov/gophermart/internal/sharedkernel"

// UserOrderNumber is now represent users registered order.
type UserOrderNumber struct {
	id      string
	user    string
	status  sharedkernel.Status
	number  int
	accrual int
}

func NewOrderNumber(number, accrual int, userID string, status sharedkernel.Status) UserOrderNumber {
	return UserOrderNumber{
		id:      sharedkernel.NewUUID(),
		user:    userID,
		number:  number,
		status:  status,
		accrual: accrual,
	}
}
