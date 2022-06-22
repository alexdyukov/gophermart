package sharedkernel

import "github.com/google/uuid"

func NewUUID() string {
	id, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	return id.String()
}
