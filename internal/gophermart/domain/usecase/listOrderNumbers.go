package usecase

import (
	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

type (
	ListOrderNumsRepository interface {
		GetOrdersByUser(string) []core.OrderNumber
	}

	ListOrderNumsInputPort interface {
		Execute(user *sharedkernel.User) ([]core.OrderNumber, error)
	}

	ListOrderNums struct {
		Repo ListOrderNumsRepository
	}
)

func NewListOrderNums(repo ListOrderNumsRepository) *ListOrderNums {
	return &ListOrderNums{
		Repo: repo,
	}
}

func (l *ListOrderNums) Execute(user *sharedkernel.User) ([]core.OrderNumber, error) {
	orders := l.Repo.GetOrdersByUser(user.ID())

	return orders, nil
}
