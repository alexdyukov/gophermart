package usecase

import (
	"context"
	"errors"

	"github.com/alexdyukov/gophermart/internal/accrual/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

type (
	ShowOrderCalculationRepository interface {
		GetOrderByNumber(context.Context, int) (*core.OrderReceipt, error)
	}

	ShowOrderCalculationPrimaryPort interface {
		Execute(context.Context, int) (*ShowOrderCalculationOutputDTO, error)
	}

	ShowOrderCalculationOutputDTO struct {
		Status  string             `json:"status"`
		Order   int                `json:"order"`
		Accrual sharedkernel.Money `json:"accrual"`
	}

	ShowOrderCalculation struct {
		Repo ShowOrderCalculationRepository
	}
)

var ErrOrderReceiptNotExist = errors.New("order receipt does not exist")

func NewShowOrderCalculation(repo ShowOrderCalculationRepository) *ShowOrderCalculation {
	return &ShowOrderCalculation{
		Repo: repo,
	}
}

func (s *ShowOrderCalculation) Execute(ctx context.Context, number int) (*ShowOrderCalculationOutputDTO, error) {
	orderState, err := s.Repo.GetOrderByNumber(ctx, number)
	if err != nil {
		return nil, err
	}

	output := ShowOrderCalculationOutputDTO{
		Order:   orderState.OrderNumber,
		Status:  orderState.Status.String(),
		Accrual: orderState.Accrual,
	}

	return &output, nil
}
