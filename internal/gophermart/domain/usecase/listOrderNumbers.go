package usecase

import "github.com/alexdyukov/gophermart/internal/gophermart/domain/core"

type ListOrderNumsRepository interface {
	GetOrdersByUser(string) []core.OrderNumber
}

type ListOrderNumsInputPort interface {
	Execute(string) ([]core.OrderNumber, error)
}

func NewListOrderNums(repo ListOrderNumsRepository) *ListOrderNums {
	return &ListOrderNums{
		Repo: repo,
	}
}

type ListOrderNums struct {
	Repo ListOrderNumsRepository
}

// constructor
// func NewListOrderNums ...

func (l *ListOrderNums) Execute(user string) ([]core.OrderNumber, error) {
	// checkings
	orders := l.Repo.GetOrdersByUser(user)
	return orders, nil
}
