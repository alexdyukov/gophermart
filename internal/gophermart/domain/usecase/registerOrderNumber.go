package usecase

import (
	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

type (

	// RegisterUserOrderRepository is a secondary port.
	RegisterUserOrderRepository interface {
		SaveUserOrder(core.UserOrderNumber) error
	}

	// CalculationStateGateway is a secondary port.
	CalculationStateGateway interface {
		GetOrderCalculationState(int) (*CalculationStateDTO, error)
	}

	// CalculationStateDTO is secondary DTO.
	CalculationStateDTO struct {
		Status  sharedkernel.Status `json:"status"`
		Order   int                 `json:"order"`
		Accrual int                 `json:"accrual"`
	}

	// RegisterUserOrderPrimaryPort is a primary port.
	RegisterUserOrderPrimaryPort interface {
		Execute(int, *sharedkernel.User) error
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

func (ruo *RegisterUserOrder) Execute(number int, user *sharedkernel.User) error {
	inputDTO, err := ruo.ServiceGateway.GetOrderCalculationState(number)
	if err != nil {
		return err // nolint:wrapcheck // ok
	}

	userOrder := core.NewOrderNumber(number, inputDTO.Accrual, user.ID(), inputDTO.Status)

	err = ruo.Repository.SaveUserOrder(userOrder)
	if err != nil {
		return err // nolint:wrapcheck // ok
	}

	return nil
}
