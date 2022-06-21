package usecase

import "github.com/alexdyukov/gophermart/internal/gophermart/domain/core"

type LoadOrderNumberRepository interface {
	SaveOrderNumber(core.OrderNumber) error
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
	orderNumber := core.NewOrderNumber(number)
	l.Repo.SaveOrderNumber(orderNumber)
	return nil
}
