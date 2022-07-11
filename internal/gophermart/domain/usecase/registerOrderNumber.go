package usecase

import (
	"context"
	"log"
	"strconv"

	"github.com/alexdyukov/gophermart/internal/accrual/domain/usecase"
	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

type (

	// RegisterUserOrderRepository is a secondary port.
	RegisterUserOrderRepository interface {
		SaveUserOrder(context.Context, *core.UserOrderNumber) error
		UpdateUserBalance(context.Context, []string) error
	}

	// CalculationStateGateway is a secondary port.
	CalculationStateGateway interface {
		GetOrderCalculationState(int64) (*CalculationStateDTO, error)
	}

	// CalculationStateDTO is secondary DTO.
	CalculationStateDTO struct {
		Order   string              `json:"order"`
		Status  sharedkernel.Status `json:"status"`
		Accrual sharedkernel.Money  `json:"accrual"`
	}

	// RegisterUserOrderPrimaryPort is a primary port.
	RegisterUserOrderPrimaryPort interface {
		Execute(context.Context, string, *sharedkernel.User) error
	}

	// RegisterUserOrder is a usecase.
	RegisterUserOrder struct {
		Repository     RegisterUserOrderRepository
		ServiceGateway CalculationStateGateway
	}
)

func NewLoadOrderNumber(repo RegisterUserOrderRepository, gw CalculationStateGateway) *RegisterUserOrder {
	return &RegisterUserOrder{
		Repository:     repo,
		ServiceGateway: gw,
	}
}

func (ruo *RegisterUserOrder) Execute(ctx context.Context, number string, user *sharedkernel.User) error {
	if !sharedkernel.ValidLuhn(number) {
		return sharedkernel.ErrIncorrectOrderNumber
	}

	orderNumber, err := strconv.ParseInt(number, 10, 64) // nolint:gomnd // ok
	if err != nil {
		return usecase.ErrIncorrectOrderNumber
	}

	inputDTO, err := ruo.ServiceGateway.GetOrderCalculationState(orderNumber)
	if err != nil {
		log.Printf("%v", err)
	}

	if inputDTO == nil {
		inputDTO = &CalculationStateDTO{
			Accrual: 0,
			Order:   number,
			Status:  sharedkernel.NEW,
		}
	}

	userOrder := core.NewOrderNumber(orderNumber, inputDTO.Accrual, user.ID(), inputDTO.Status)

	err = ruo.Repository.SaveUserOrder(ctx, &userOrder)
	if err != nil {
		return err // nolint:wrapcheck // ok
	}

	if inputDTO.Status == sharedkernel.PROCESSED {
		sliceUsers := make([]string, 0)
		sliceUsers = append(sliceUsers, user.ID())
		err = ruo.Repository.UpdateUserBalance(ctx, sliceUsers)

		if err != nil {
			return err // nolint:wrapcheck // ok
		}
	}

	return nil
}
