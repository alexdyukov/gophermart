package sharedkernel

import (
	"errors"
	"fmt"
)

type AppError struct {
	err error
	// packageName string.
	// funcName    string.
}

var errNotValOrdNum = errors.New("repositoryError")

func (e *AppError) Error() string {
	return e.err.Error()
}

func ErrorNotValidOrderNumber(msg string) error {
	return fmt.Errorf("%w: %s", errNotValOrdNum, msg)
}
