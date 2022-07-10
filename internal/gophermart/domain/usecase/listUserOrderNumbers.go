package usecase

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

type (
	ListUserOrdersRepository interface {
		FindAllOrders(context.Context, string) ([]core.UserOrderNumber, error)
		SaveUserOrder(context.Context, *core.UserOrderNumber) error
	}

	ListUserOrdersPrimaryPort interface {
		Execute(context.Context, *sharedkernel.User) ([]ListUserOrdersOutputDTO, error)
	}

	ListCalculationStateGateway interface {
		GetOrderCalculationState(int64) (*CalculationStateDTO, error)
	}

	ListUserOrdersOutputDTO struct {
		UploadedAt    time.Time          `json:"-"`
		UploadedAtStr string             `json:"uploaded_at"` // nolint:tagliatelle // ok
		Number        string             `json:"number"`
		Status        string             `json:"status"`
		Accrual       sharedkernel.Money `json:"accrual"`
	}

	ListUserOrders struct {
		Repo           ListUserOrdersRepository
		ServiceGateway ListCalculationStateGateway
	}
)

func NewListUserOrders(repo ListUserOrdersRepository, gw ListCalculationStateGateway) *ListUserOrders {
	return &ListUserOrders{
		Repo:           repo,
		ServiceGateway: gw,
	}
}

func (lou *ListUserOrders) Execute(ctx context.Context, user *sharedkernel.User) ([]ListUserOrdersOutputDTO, error) {
	// Update orders
	ord, err := lou.Repo.FindAllOrders(ctx, user.ID())
	if err != nil {
		return nil, err
	}

	for _, order := range ord {
		inputDTO, err := lou.ServiceGateway.GetOrderCalculationState(order.Number) // nolint:govet // ok.
		if err != nil {
			continue
		}

		if inputDTO == nil {
			continue
		}

		userOrder := core.NewOrderNumber(order.Number, inputDTO.Accrual, user.ID(), inputDTO.Status)

		err = lou.Repo.SaveUserOrder(ctx, &userOrder)
		if err != nil {
			continue
		}
	}

	orders, err := lou.Repo.FindAllOrders(ctx, user.ID())
	if err != nil {
		return nil, err
	}

	log.Println(orders)

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
