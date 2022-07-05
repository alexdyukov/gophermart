package usecase

import (
	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

type (

	// RegisterOrderRepository is a secondary port
	RegisterOrderRepository interface {
		SaveOrderNumber(core.OrderNumber) error
	}

	// RegisterOrderCalculationStateGateway is a secondary port
	RegisterOrderCalculationStateGateway interface {
		GetOrderCalculationState(int) (*RegisterOrderCalculationStateDTO, error)
	}

	// RegisterOrderCalculationStateDTO is secondary DTO
	RegisterOrderCalculationStateDTO struct {
		Order   string              `json:"order"`
		Status  sharedkernel.Status `json:"status"`
		Accrual int                 `json:"accrual"`
	}

	// RegisterOrderPrimaryPort is a primary port
	RegisterOrderPrimaryPort interface {
		Execute(int, *sharedkernel.User) error
	}

	// RegisterOrder is a usecase
	RegisterOrder struct {
		Repository     RegisterOrderRepository
		ServiceGateway RegisterOrderCalculationStateGateway
	}
)

func NewLoadOrderNumber(repo RegisterOrderRepository, gw RegisterOrderCalculationStateGateway) *RegisterOrder {
	return &RegisterOrder{
		Repository:     repo,
		ServiceGateway: gw,
	}
}

func (ro *RegisterOrder) Execute(number int, user *sharedkernel.User) error {

	// todo: check incoming number

	inputDTO, err := ro.ServiceGateway.GetOrderCalculationState(number)
	if err != nil {
		return err //nolint:wrapcheck // ok
	}

	orderNumber := core.NewOrderNumber(number, inputDTO.Accrual, user.ID(), inputDTO.Status)

	err = ro.Repository.SaveOrderNumber(orderNumber)
	if err != nil {
		return err //nolint:wrapcheck // ok
	}

	return nil
}
