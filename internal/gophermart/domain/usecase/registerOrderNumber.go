package usecase

import (
	"context"
	"fmt"
	"strconv"

	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

type (

	// RegisterUserOrderRepository is a secondary port.
	RegisterUserOrderRepository interface {
		SaveUserOrder(context.Context, *core.UserOrderNumber) error
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
		Execute(context.Context, int, *sharedkernel.User) error
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

func (ruo *RegisterUserOrder) Execute(ctx context.Context, number int, user *sharedkernel.User) error {

	if !sharedkernel.ValidLuhn(strconv.Itoa(number)) {
		fmt.Println("#RegisterUserOrderPostHandler: номер не прошел проверку луна ", number)
		return sharedkernel.ErrIncorrectOrderNumber
	}

	// вот тут выдает ошибку, не может что-то там обновить и дальше не идет.
	inputDTO, err := ruo.ServiceGateway.GetOrderCalculationState(number)
	fmt.Println("#RegisterUserOrderPostHandler: из аккурала получили такую структуру: ", inputDTO)
	if err != nil {

		// return err // nolint:wrapcheck // ok
	}

	userOrder := core.NewOrderNumber(number, sharedkernel.Money(inputDTO.Accrual), user.ID(), inputDTO.Status)

	fmt.Println("#RegisterUserOrderPostHandler: номер верный идем дальше, пробуем записать номер в БД ", number)
	err = ruo.Repository.SaveUserOrder(ctx, &userOrder)
	if err != nil {
		fmt.Println("#RegisterUserOrderPostHandler:", err)
		return err // nolint:wrapcheck // ok
	}

	fmt.Println("#RegisterUserOrderPostHandler: заказ в БД записан УСПЕШНО ", number)
	return nil
}
