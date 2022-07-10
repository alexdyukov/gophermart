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

	// CalculationStateDTO is secondary DTO
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

func (l *ListUserOrders) Execute(ctx context.Context, user *sharedkernel.User) ([]ListUserOrdersOutputDTO, error) {
	// Получили все заказы, пробуем обновить их статусы.
	orders, err := l.Repo.FindAllOrders(ctx, user.ID())
	if err != nil {
		return nil, err
	}

	lstOrdNumsDTO := make([]ListUserOrdersOutputDTO, 0, len(orders))

	for _, order := range orders {
		log.Printf("#ListUserOrdersGetHandler пробуем отправлять номер заказа %v для проверки в accrual ", order.Number)
		inputDTO, err := l.ServiceGateway.GetOrderCalculationState(order.Number)
		if err != nil {
			log.Printf("GetOrderCalculationState получили ошибку %v", err)
			log.Printf("%v", err)
		}

		if inputDTO == nil {
			log.Printf("#ListUserOrdersGetHandler пустая inputDTO, выполняем continue")
			continue
		}

		log.Printf("#ListUserOrdersGetHandler: получили данные из accrual по переданному заказу: ", inputDTO)

		userOrder := core.NewOrderNumber(order.Number, inputDTO.Accrual, user.ID(), inputDTO.Status)
		log.Println("#RegisterUserOrderPostHandler вот такая структура userOrder:", userOrder)

		err = l.Repo.SaveUserOrder(ctx, &userOrder)
		if err != nil {
			log.Printf("ListUserOrdersGetHandler вышла ошибка при обовлении данных по заказу в БД", err)
			continue
		}
	}

	//это функция была...
	orders, err = l.Repo.FindAllOrders(ctx, user.ID())
	if err != nil {
		return nil, err
	}

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
