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
