package usecase

import (
	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

type LoadOrderNumberRepository interface {
	SaveOrderNumber(core.UserOrderNumber) error
}

type LoadOrderNumberInputPort interface {
	Execute(int) error
}

func NewLoadOrderNumber(repo LoadOrderNumberRepository) *LoadOrderNumber {
	return &LoadOrderNumber{
		Repo: repo,
	}
}

type LoadOrderNumber struct {
	Repo LoadOrderNumberRepository
}

func (l LoadOrderNumber) Execute(number int) error {
	// do work...
	orderNumber := core.NewOrderNumber(number, 33333, "werwerwer", sharedkernel.NEW)
	l.Repo.SaveOrderNumber(orderNumber)
	return nil
}
