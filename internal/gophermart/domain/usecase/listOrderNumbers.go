package usecase

import (
	"context"
	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
	"sort"
	"time"
)

type ListOrderNumsRepository interface {
	GetOrdersByUser(context.Context, string) ([]core.OrderNumber, error)
}

type ListOrderNumsInputPort interface {
	Execute(context.Context, string) ([]ListOrderNumsDTO, error)
}

func NewListOrderNums(repo ListOrderNumsRepository) *ListOrderNums {
	return &ListOrderNums{
		Repo: repo,
	}
}

type ListOrderNums struct {
	Repo ListOrderNumsRepository
}

// constructor
// func NewListOrderNums ...

type ListOrderNumsDTO struct {
	Number  string    `json:"number"`
	Status  string    `json:"status"`
	Accrual float32   `json:"accrual,omitempty"`
	Data    time.Time `json:"-"`
	DataStr string    `json:"uploaded_at"`
}

func (l *ListOrderNums) Execute(ctx context.Context, user string) ([]ListOrderNumsDTO, error) {
	// checkings
	orders, err := l.Repo.GetOrdersByUser(ctx, user)

	if err != nil {
		return nil, err
	}

	lstOrdNumsDTO := make([]ListOrderNumsDTO, 0)

	for _, order := range orders {
		lstOrdNumsDTO = append(lstOrdNumsDTO, ListOrderNumsDTO{Number: order.Number, Status: order.Status.String(), Accrual: order.Accrual, Data: order.Data, DataStr: order.Data.Format(time.RFC3339)})
	}

	sort.SliceStable(lstOrdNumsDTO, func(i, j int) bool {
		return lstOrdNumsDTO[i].Data.Before(lstOrdNumsDTO[j].Data)
	})

	return lstOrdNumsDTO, nil
}
