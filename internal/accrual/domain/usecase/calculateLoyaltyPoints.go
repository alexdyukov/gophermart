package usecase

type (
	CalculateLoyaltyPointsRepository interface {
		SaveCalculationProcessData() error
	}

	CalculateLoyaltyPointsInputDTO struct {
		/* needed data */
	}

	CalculateLoyaltyPointsInputPort interface {
		Execute(CalculateLoyaltyPointsInputDTO) error
	}

	CalculateLoyaltyPoints struct {
		repo CalculateLoyaltyPointsRepository
	}
)

func NewCalculateLoyaltyPoints(repo CalculateLoyaltyPointsRepository) *CalculateLoyaltyPoints {
	return &CalculateLoyaltyPoints{
		repo: repo,
	}
}

func (c *CalculateLoyaltyPoints) Execute(CalculateLoyaltyPointsInputDTO) error {
	// i.e. save to DB calculation data to calculate points
	// in async approach
	return nil
}
