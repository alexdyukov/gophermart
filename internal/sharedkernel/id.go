package sharedkernel

import "github.com/google/uuid"

// NewUUID is UUID generator Facade.
func NewUUID() string {
	uid, err := uuid.NewRandom()
	if err != nil {
		panic(err) // do not panic, handle in another way
	}

	return uid.String()
}
