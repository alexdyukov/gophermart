package usecase

import (
	"context"
	"errors"
	"strconv"

	"github.com/alexdyukov/gophermart/internal/accrual/domain/core"
)

type (
	RegisterOrderReceiptRepository interface {
		SaveOrderReceipt(context.Context, *core.OrderReceipt) error
	}

	RegisterOrderReceiptPrimaryPort interface {
		Execute(context.Context, *RegisterOrderReceiptInputDTO) (*core.OrderReceipt, error)
	}

	RegisterOrderReceiptInputDTO struct {
		OrderNumber string         `json:"order"`
		Goods       []core.Product `json:"goods"`
	}

	RegisterOrderReceipt struct {
		repo RegisterOrderReceiptRepository
	}
)

var ErrOrderAlreadyExist = errors.New("error order number already exists")

func NewRegisterOrderReceipt(repo RegisterOrderReceiptRepository) *RegisterOrderReceipt {
	return &RegisterOrderReceipt{
		repo: repo,
	}
}

func (reg *RegisterOrderReceipt) Execute(
	ctx context.Context, dto *RegisterOrderReceiptInputDTO,
) (*core.OrderReceipt, error) { // nolint:whitespace // ok
	number, err := strconv.Atoi(dto.OrderNumber)
	if err != nil {
		return nil, err
	}

	orderReceipt := core.NewOrderReceipt(number, dto.Goods)

	err = reg.repo.SaveOrderReceipt(ctx, orderReceipt)
	if err != nil {
		return nil, err
	}

	// start sync or async calculation...

	return orderReceipt, nil
}
