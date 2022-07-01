package usecase

import "github.com/alexdyukov/gophermart/internal/gophermart/domain/core"

type (
	LoadOrderNumberRepository interface {
		SaveOrderNumber(core.OrderNumber) error
	}

	LoadOrderNumberInputPort interface {
		Execute(int) error
	}

	LoadOrderNumber struct {
		Repo LoadOrderNumberRepository
	}
)

func NewLoadOrderNumber(repo LoadOrderNumberRepository) *LoadOrderNumber {
	return &LoadOrderNumber{
		Repo: repo,
	}
}

func (l LoadOrderNumber) Execute(number int) error {
	orderNumber := core.NewOrderNumber(number)

	err := l.Repo.SaveOrderNumber(orderNumber)
	if err != nil {
		return err //nolint:wrapcheck // ok
	}

	return nil
}
