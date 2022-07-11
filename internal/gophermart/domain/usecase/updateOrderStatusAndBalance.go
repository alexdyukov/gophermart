package usecase

import (
	"context"
	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
	"log"
)

type (
	UpdateUserOrderBalanceRepository interface {
		FindAllUnprocessedOrders(context.Context) ([]core.UserOrderNumber, error)
		SaveOrderWithoutCheck(context.Context, *core.UserOrderNumber) error
		UpdateUserBalance(context.Context, []string) error
	}

	UpdateUsrOrderAndBalancePrimaryPort interface {
		Execute(context.Context) error
	}

	UpdateCalculationStateGateway interface {
		GetOrderCalculationState(int64) (*CalculationStateDTO, error)
	}

	UpdateOrderAndBalance struct {
		Repo           UpdateUserOrderBalanceRepository
		ServiceGateway UpdateCalculationStateGateway
	}
)

func NewUpdateOrderAndBalance(repo UpdateUserOrderBalanceRepository, gw UpdateCalculationStateGateway) *UpdateOrderAndBalance {
	return &UpdateOrderAndBalance{
		Repo:           repo,
		ServiceGateway: gw,
	}
}

func (uob *UpdateOrderAndBalance) Execute(ctx context.Context) error {

	allOrders, err := uob.Repo.FindAllUnprocessedOrders(ctx)
	if err != nil {
		log.Println("UpdateOrderAndBalance #1:FindAllUnprocessedOrders - ошибка получения всех заказов")
		return err // nolint:wrapcheck // ok
	}

	sliceUsers := make([]string, 0)

	for _, order := range allOrders {
		inputDTO, err := uob.ServiceGateway.GetOrderCalculationState(order.Number) // nolint:govet // ok.
		if err != nil {
			log.Println("UpdateOrderAndBalance #1: ошибка получения GetOrderCalculationState")
			continue
		}

		if inputDTO == nil {
			continue
		}

		userOrder := core.NewOrderNumber(order.Number, inputDTO.Accrual, order.User, inputDTO.Status)

		if inputDTO.Status != order.Status {
			sliceUsers = append(sliceUsers, order.User)
			err = uob.Repo.SaveOrderWithoutCheck(ctx, &userOrder)
			if err != nil {
				log.Println("UpdateOrderAndBalance #1: ошибка сохранения заказа в бд")
				continue
			}
		}
	}

	sliceUsers = removeDuplicateElement(sliceUsers)

	if len(sliceUsers) > 0 {
		err = uob.Repo.UpdateUserBalance(ctx, sliceUsers)
	}
	return nil
}

func removeDuplicateElement(sliceEl []string) []string {
	result := make([]string, 0, len(sliceEl))
	temp := map[string]struct{}{}
	for _, item := range sliceEl {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
