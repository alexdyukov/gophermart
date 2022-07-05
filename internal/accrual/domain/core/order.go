package core

import (
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

type (
	// Product should not contain any tags, violate restrictions
	// for sake of simplicity as this is study project.
	Product struct {
		Description string             `json:"description"`
		Price       sharedkernel.Money `json:"price"`
	}

	OrderReceipt struct {
		Goods       []Product
		Accrual     int
		OrderNumber int
		Status      sharedkernel.Status
	}
)

func NewOrderReceipt(number int, goods []Product) *OrderReceipt {
	order := OrderReceipt{
		Status:      sharedkernel.NEW,
		Accrual:     0,
		OrderNumber: number,
		Goods:       goods,
	}

	return &order
}

func (o *OrderReceipt) CalculateRewardPoints(rewards []Reward) {
	// Temporary here. Chances are will move into separate domain entity
}
