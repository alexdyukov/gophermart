package usecase

import (
	"log"
	"time"

	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

type (
	ListUserOrdersRepository interface {
		GetOrdersByUser(string) []core.OrderNumber
	}

	ListUserOrdersInputPort interface {
		Execute(user *sharedkernel.User) ([]ListUserOrdersOutputDTO, error)
	}

	ListUserOrdersOutputDTO struct {
		UploadedAt time.Time `json:"uploaded_at"` // nolint:tagliatelle // ok
		Number     string    `json:"number"`
		Status     string    `json:"status"`
		Accrual    int       `json:"accrual"`
	}

	ListUserOrders struct {
		Repo ListUserOrdersRepository
	}
)

func NewListOrderNums(repo ListUserOrdersRepository) *ListUserOrders {
	return &ListUserOrders{
		Repo: repo,
	}
}

func (l *ListUserOrders) Execute(user *sharedkernel.User) ([]ListUserOrdersOutputDTO, error) {
	orders := l.Repo.GetOrdersByUser(user.ID())

	log.Println(orders)

	return []ListUserOrdersOutputDTO{
		{},
	}, nil
}
