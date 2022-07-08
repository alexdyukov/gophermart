package usecase

import (
	"context"
	"errors"
	"strconv"

	"github.com/alexdyukov/gophermart/internal/accrual/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

type (
	ShowOrderCalculationRepository interface {
		GetOrderByNumber(context.Context, int64) (*core.OrderReceipt, error)
	}

	ShowOrderCalculationPrimaryPort interface {
		Execute(context.Context, string) (*ShowOrderCalculationOutputDTO, error)
	}

	ShowOrderCalculationOutputDTO struct {
		Status  string             `json:"status"`
		Order   int64              `json:"order"`
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

func (s *ShowOrderCalculation) Execute(ctx context.Context, number string) (*ShowOrderCalculationOutputDTO, error) {
	if !sharedkernel.ValidLuhn(number) {
		return nil, ErrIncorrectOrderNumber
	}

	orderNumber, err := strconv.ParseInt(number, 10, 64) // nolint:gomnd // ok
	if err != nil {
		return nil, ErrIncorrectOrderNumber
	}

	orderState, err := s.Repo.GetOrderByNumber(ctx, orderNumber)
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
