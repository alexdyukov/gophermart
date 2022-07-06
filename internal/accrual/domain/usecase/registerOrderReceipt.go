package usecase

import (
	"github.com/alexdyukov/gophermart/internal/accrual/domain/core"
)

type (
	RegisterOrderReceiptRepository interface {
		SaveOrderReceipt(*core.OrderReceipt) error
	}

	RegisterOrderReceiptPrimaryPort interface {
		Execute(receipt RegisterOrderReceiptInputDTO) error
	}

	RegisterOrderReceiptInputDTO struct {
		Goods       []core.Product `json:"goods"`
		OrderNumber int            `json:"order"`
	}

	RegisterOrderReceipt struct {
		repo RegisterOrderReceiptRepository
	}
)

func NewRegisterOrderReceipt(repo RegisterOrderReceiptRepository) *RegisterOrderReceipt {
	return &RegisterOrderReceipt{
		repo: repo,
	}
}

func (c *RegisterOrderReceipt) Execute(dto RegisterOrderReceiptInputDTO) error {
	orderReceipt := core.NewOrderReceipt(dto.OrderNumber, dto.Goods)

	err := c.repo.SaveOrderReceipt(orderReceipt)
	if err != nil {
		return err
	}

	// start sync or async calculation...

	return nil
}
