package usecase

// async usecase

type CalculateLoyaltyPointsRepository interface {
	SaveCalculationProcessData() error
}

type CalculateLoyaltyPointsInputDTO struct {
	/* needed data */
}

type CalculateLoyaltyPointsInputPort interface {
	Execute(CalculateLoyaltyPointsInputDTO) error
}

type CalculateLoyaltyPoints struct {
	repo CalculateLoyaltyPointsRepository
}

func NewCalculateLoyaltyPoints(repo CalculateLoyaltyPointsRepository) *CalculateLoyaltyPoints {
	return &CalculateLoyaltyPoints{
		repo: repo,
	}
}

func (c *CalculateLoyaltyPoints) Execute(dto CalculateLoyaltyPointsInputDTO) error {
	// todo: make needed checks
	// todo: start calculation process
	// i.e. save to DB calculation data to calculate points
	// in async approach
	return nil
}
