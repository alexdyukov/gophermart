package usecase

import "context"

// async usecase

type CalculateLoyaltyPointsRepository interface {
	SaveCalculationProcessData() error
}

type CalculateLoyaltyPointsInputDTO struct {
	/* needed data */
}

type CalculateLoyaltyPointsInputPort interface {
	Execute(context.Context, CalculateLoyaltyPointsInputDTO) error
}

type CalculateLoyaltyPoints struct {
	repo CalculateLoyaltyPointsRepository
}

func NewCalculateLoyaltyPoints(repo CalculateLoyaltyPointsRepository) *CalculateLoyaltyPoints {
	return &CalculateLoyaltyPoints{
		repo: repo,
	}
}

func (c *CalculateLoyaltyPoints) Execute(ctx context.Context, dto CalculateLoyaltyPointsInputDTO) error {
	// todo: make needed checks
	// todo: start calculation process
	// i.e. save to DB calculation data to calculate points
	// in async approach
	return nil
}
