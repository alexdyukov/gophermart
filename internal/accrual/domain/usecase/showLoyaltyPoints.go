package usecase

type (
	ShowOrderCalculationRepository interface {
		GetLoyaltyPointsByOrderNumber(int) error
	}

	ShowOrderCalculationPrimaryPort interface {
		Execute(int) error
	}

	ShowOrderCalculation struct {
		Repo ShowOrderCalculationRepository
	}
)

func NewShowLoyaltyPoints(repo ShowOrderCalculationRepository) *ShowOrderCalculation {
	return &ShowOrderCalculation{
		Repo: repo,
	}
}

func (s *ShowOrderCalculation) Execute(number int) error {
	err := s.Repo.GetLoyaltyPointsByOrderNumber(number)
	if err != nil {
		return err //nolint:wrapcheck // ok
	}

	return nil
}
