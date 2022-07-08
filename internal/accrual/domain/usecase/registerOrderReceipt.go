package usecase

import (
	"context"
	"errors"
	"strconv"

	"github.com/alexdyukov/gophermart/internal/accrual/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
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

var (
	ErrOrderAlreadyExist    = errors.New("error order number already exists")
	ErrIncorrectOrderNumber = errors.New("order number is incorrect")
)

func NewRegisterOrderReceipt(repo RegisterOrderReceiptRepository) *RegisterOrderReceipt {
	return &RegisterOrderReceipt{
		repo: repo,
	}
}

func (reg *RegisterOrderReceipt) Execute(
	ctx context.Context, dto *RegisterOrderReceiptInputDTO,
) (*core.OrderReceipt, error) { // nolint:whitespace // ok
	if !sharedkernel.ValidLuhn(dto.OrderNumber) {
		return nil, ErrIncorrectOrderNumber
	}

	number, err := strconv.ParseInt(dto.OrderNumber, 10, 64) // nolint:gomnd // ok
	if err != nil {
		return nil, ErrIncorrectOrderNumber
	}

	orderReceipt := core.NewOrderReceipt(number, dto.Goods)

	err = reg.repo.SaveOrderReceipt(ctx, orderReceipt)
	if err != nil {
		return nil, err
	}

	// start sync or async calculation...

	return orderReceipt, nil
}
