package usecase

type (
	RegisterPurchasedOrderRepository interface {
		SavePurchasedOrder() error
	}

	RegisterPurchasedOrderInputDTO struct {
		/* needed data */
	}

	RegisterPurchasedOrderPrimaryPort interface {
		Execute(RegisterPurchasedOrderInputDTO) error
	}

	RegisterPurchasedOrder struct {
		repo RegisterPurchasedOrderRepository
	}
)

func NewCalculateLoyaltyPoints(repo RegisterPurchasedOrderRepository) *RegisterPurchasedOrder {
	return &RegisterPurchasedOrder{
		repo: repo,
	}
}

func (c *RegisterPurchasedOrder) Execute(RegisterPurchasedOrderInputDTO) error {
	return nil
}
