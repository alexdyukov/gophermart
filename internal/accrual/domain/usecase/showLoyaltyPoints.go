package usecase

type ShowLoyaltyPointsRepository interface {
	GetLoyaltyPointsByOrderNumber(int) error
}

type ShowLoyaltyPointsInputPort interface {
	Execute(int) error
}

type ShowLoyaltyPoints struct {
	Repo ShowLoyaltyPointsRepository
}

func NewShowLoyaltyPoints(repo ShowLoyaltyPointsRepository) *ShowLoyaltyPoints {
	return &ShowLoyaltyPoints{
		Repo: repo,
	}
}

func (s *ShowLoyaltyPoints) Execute(number int) error {
	s.Repo.GetLoyaltyPointsByOrderNumber(number)
	return nil
}
