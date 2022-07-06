package usecase

import (
	"github.com/alexdyukov/gophermart/internal/accrual/domain/core"
)

type (
	ShowOrderCalculationRepository interface {
		GetOrderByNumber(int) (core.OrderReceipt, error)
	}

	ShowOrderCalculationPrimaryPort interface {
		Execute(int) (*ShowOrderCalculationOutputDTO, error)
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

func (s *ShowOrderCalculation) Execute(number int) (*ShowOrderCalculationOutputDTO, error) {
	orderState, err := s.Repo.GetOrderByNumber(number)
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
