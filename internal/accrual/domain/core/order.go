package core

import (
	"fmt"
	"time"

	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

type (
	// Product should not contain any tags, violate restrictions
	// for sake of simplicity as this is study project.
	Product struct {
		Match       string
		Description string             `json:"description"`
		Price       sharedkernel.Money `json:"price"`
	}

	OrderReceipt struct {
		Goods       []Product
		Accrual     sharedkernel.Money
		OrderNumber int64
		Status      sharedkernel.Status
	}
)

func NewOrderReceipt(number int64, goods []Product) *OrderReceipt {
	order := OrderReceipt{
		Status:      sharedkernel.NEW,
		Accrual:     0,
		OrderNumber: number,
		Goods:       goods,
	}

	return &order
}

func (ord *OrderReceipt) CalculateRewardPoints(rewards map[string]Reward) {
	time.Sleep(time.Second * 1) // nolint:gomnd // temp sleep

	points := sharedkernel.Money(0)

	for _, v := range ord.Goods {
		rew := rewards[v.Match]
		if rew.isPercentage() {

			percentPoints := (v.Price / 100) * rew.RewardPoints() // nolint:gomnd // percent number
			fmt.Println("PERCENT reward points->", percentPoints)

			points += percentPoints

			continue
		}

		fmt.Println("PT reward points->", rew.RewardPoints(), rew.Match())
		points += rew.RewardPoints()
	}

	ord.Accrual = points
	ord.Status = sharedkernel.PROCESSED
}
