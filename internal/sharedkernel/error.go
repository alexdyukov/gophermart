package sharedkernel

import (
	"errors"
)

type AppError struct {
	err error
	// packageName string.
	// funcName    string.
}

var (
	ErrOrderExists          = errors.New("order exists")
	ErrAnotherUserOrder     = errors.New("this is anoter user's order")
	ErrIncorrectOrderNumber = errors.New("incorrect order number")
	ErrInsufficientFunds    = errors.New("insufficient funds")
)

func (e *AppError) Error() string {
	return e.err.Error()
}
