package usecase

import (
	"context"
	"strconv"

	"github.com/alexdyukov/gophermart/internal/accrual/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

type (
	ShowOrderCalculationRepository interface {
		GetOrderByNumber(context.Context, int) (core.OrderReceipt, error)
	}

	ShowOrderCalculationPrimaryPort interface {
		Execute(context.Context, int) (*ShowOrderCalculationOutputDTO, error)
	}

	ShowOrderCalculationOutputDTO struct {
		Status  string `json:"status"`
		Order   int    `json:"order"`
		Accrual int    `json:"accrual"`
	}

	ShowOrderCalculation struct {
		Repo ShowOrderCalculationRepository
	}
)

func NewShowOrderCalculation(repo ShowOrderCalculationRepository) *ShowOrderCalculation {
	return &ShowOrderCalculation{
		Repo: repo,
	}
}

func (s *ShowOrderCalculation) Execute(ctx context.Context, number int) (*ShowOrderCalculationOutputDTO, error) {
	if sharedkernel.ValidLuhn(strconv.Itoa(number)) {
		orderState, err := s.Repo.GetOrderByNumber(ctx, number)
		if err != nil {
			return nil, err //nolint:wrapcheck // ok
		}

		output := ShowOrderCalculationOutputDTO{
			Order:   orderState.OrderNumber,
			Status:  orderState.Status.String(),
			Accrual: orderState.Accrual,
		}

		return &output, nil
	}

	return &ShowOrderCalculationOutputDTO{}, sharedkernel.ErrorNotValidOrderNumber(strconv.Itoa(number))
}
