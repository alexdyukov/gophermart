package usecase

import "github.com/alexdyukov/gophermart/internal/gophermart/domain/core"

type ShowBalanceStateRepo interface {
	GetAccountByID(string) (core.Account, error)
}

type ShowBalanceStateInputPort interface {
	Execute(string) (ShowBalanceStateOutputDTO, error)
}

type ShowBalanceStateOutputDTO struct {
	// current..
	// withdrawn..
}

type ShowBalanceState struct {
	Repo ShowBalanceStateRepo
}

func NewShowBalanceState(repo ShowBalanceStateRepo) *ShowBalanceState {
	return &ShowBalanceState{
		Repo: repo,
	}
}

func (s *ShowBalanceState) Execute(id string) (ShowBalanceStateOutputDTO, error) {
	// checks..
	_, err := s.Repo.GetAccountByID(id)
	if err != nil {
		// process error
	}
	// construct output
	return ShowBalanceStateOutputDTO{ /* fill with data */ }, nil
}
