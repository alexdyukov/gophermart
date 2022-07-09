package sharedkernel

import (
	"errors"
	"fmt"
)

type Status int

const (
	_ Status = iota
	NEW
	PROCESSING
	INVALID
	PROCESSED
)

var ErrBadStatus = errors.New("wrong status, not supported")
var ErrOrderExists = errors.New("Order exists")
var ErrAnotherUserOrder = errors.New("This is anoter user's order")
var ErrIncorrectOrderNumber = errors.New("Incorrect order number ")

func (status Status) String() string {
	return [...]string{"NEW", "PROCESSING", "INVALID", "PROCESSED"}[status-1]
}

func (status *Status) UnmarshalJSON(bytes []byte) error {
	switch string(bytes) {
	case `"NEW"`:
		*status = NEW
	case `"PROCESSING"`:
		*status = PROCESSING
	case `"INVALID"`:
		*status = INVALID
	case `"PROCESSED"`:
		*status = PROCESSED
	default:
		return fmt.Errorf("%v %w", bytes, ErrBadStatus)
	}

	return nil
}

func (status Status) MarshalJSON() ([]byte, error) {
	return []byte(`"` + status.String() + `"`), nil
}
