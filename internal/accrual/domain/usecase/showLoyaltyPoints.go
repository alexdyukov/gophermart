package usecase

type (
	ShowLoyaltyPointsRepository interface {
		GetLoyaltyPointsByOrderNumber(int) error
	}

	ShowLoyaltyPointsInputPort interface {
		Execute(int) error
	}

	ShowLoyaltyPoints struct {
		Repo ShowLoyaltyPointsRepository
	}
)

func NewShowLoyaltyPoints(repo ShowLoyaltyPointsRepository) *ShowLoyaltyPoints {
	return &ShowLoyaltyPoints{
		Repo: repo,
	}
}

func (s *ShowLoyaltyPoints) Execute(number int) error {
	err := s.Repo.GetLoyaltyPointsByOrderNumber(number)
	if err != nil {
		return err //nolint:wrapcheck // ok
	}

	return nil
}
