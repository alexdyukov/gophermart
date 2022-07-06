package usecase

import (
	"log"
	"time"

	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

type (
	ListUserOrdersRepository interface {
		FindAllOrders(string) ([]core.UserOrderNumber, error)
	}

	ListUserOrdersPrimaryPort interface {
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

func NewListUserOrders(repo ListUserOrdersRepository) *ListUserOrders {
	return &ListUserOrders{
		Repo: repo,
	}
}

func (l *ListUserOrders) Execute(user *sharedkernel.User) ([]ListUserOrdersOutputDTO, error) {
	orders, err := l.Repo.FindAllOrders(user.ID())
	if err != nil {
		return nil, err
	}

	log.Println(orders)

	return []ListUserOrdersOutputDTO{
		{},
	}, nil
}
