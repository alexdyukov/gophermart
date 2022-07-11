package usecase

import (
	"context"
	"fmt"
	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
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
		log.Println("UpdateOrderAndBalance #1: ошибка получения всех заказов")
		return err // nolint:wrapcheck // ok
	}

	log.Println("allOrders = ", allOrders)
	sliceUsers := make([]string, 0)

	for _, order := range allOrders {
		fmt.Printf("id = %v, num = %v, usr = %v, st =%v, acc= %v \n", order.ID, order.Number, order.User, order.Status, order.Accrual)

		//inputDTO, err := uob.ServiceGateway.GetOrderCalculationState(order.Number) // nolint:govet // ok.
		err = nil
		log.Println("order = ", order)
		inputDTO := &CalculationStateDTO{
			"12345678903",
			sharedkernel.PROCESSED,
			50,
		}

		log.Println("inputDTO = ", inputDTO)
		if err != nil {
			log.Println("UpdateOrderAndBalance #1: ошибка получения GetOrderCalculationState:", err)
			continue
		}

		if inputDTO == nil {
			continue
		}

		fmt.Println("#order.User = ", order.User)
		userOrder := core.NewOrderNumber(order.Number, inputDTO.Accrual, order.User, inputDTO.Status)
		fmt.Println("#userOrder = ", userOrder.User)

		//if inputDTO.Status != order.Status {
		log.Println("UpdateOrderAndBalance сохраняем заказ")
		sliceUsers = append(sliceUsers, order.User)
		err = uob.Repo.SaveOrderWithoutCheck(ctx, &userOrder)
		if err != nil {
			log.Println("UpdateOrderAndBalance #1: ошибка сохранения заказа в бд", err)
			continue
		}
		log.Println("UpdateOrderAndBalance #1: сохранили заказ идем дальше")
		//}
	}

	sliceUsers = removeDuplicateElement(sliceUsers)

	if len(sliceUsers) > 0 {
		log.Println("UpdateOrderAndBalance #1: пробуем обновить баланс у пользователей")
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
