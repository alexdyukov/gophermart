package sharedkernel

import "errors"

type AppError struct {
	err error
	// packageName string.
	// funcName    string.
}

var (
	ErrOrderExists          = errors.New("Order exists")
	ErrAnotherUserOrder     = errors.New("This is anoter user's order")
	ErrIncorrectOrderNumber = errors.New("Incorrect order number ")
)

func (e *AppError) Error() string {
	return e.err.Error()
}
