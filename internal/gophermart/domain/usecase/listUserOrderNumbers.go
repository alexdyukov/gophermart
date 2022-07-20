package usecase

import (
	"context"
	"strconv"
	"time"

	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

type (
	ListUserOrdersRepository interface {
		FindAllOrders(context.Context, string) ([]core.UserOrderNumber, error)
	}

	ListUserOrdersPrimaryPort interface {
		Execute(context.Context, *sharedkernel.User) ([]ListUserOrdersOutputDTO, error)
	}

	ListUserOrdersOutputDTO struct {
		UploadedAt    time.Time          `json:"-"`
		UploadedAtStr string             `json:"uploaded_at"` // nolint:tagliatelle // ok
		Number        string             `json:"number"`
		Status        string             `json:"status"`
		Accrual       sharedkernel.Money `json:"accrual"`
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

func (lou *ListUserOrders) Execute(ctx context.Context, user *sharedkernel.User) ([]ListUserOrdersOutputDTO, error) {
	orders, err := lou.Repo.FindAllOrders(ctx, user.ID())
	if err != nil {
		return nil, err
	}

	lstOrdNumsDTO := make([]ListUserOrdersOutputDTO, 0, len(orders))

	for _, order := range orders {
		lstOrdNumsDTO = append(lstOrdNumsDTO, ListUserOrdersOutputDTO{
			Number:        strconv.FormatInt(order.Number, 10),
			Status:        order.Status.String(),
			Accrual:       order.Accrual,
			UploadedAt:    order.DateAndTime,
			UploadedAtStr: order.DateAndTime.Format(time.RFC3339),
		})
	}

	return lstOrdNumsDTO, nil
}
