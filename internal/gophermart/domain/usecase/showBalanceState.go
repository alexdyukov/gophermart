package usecase

import (
	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

type (
	ShowBalanceStateRepo interface {
		GetAccountByID(string) (core.Account, error)
	}

	ShowBalanceStateInputPort interface {
		Execute(user *sharedkernel.User) (*ShowBalanceStateOutputDTO, error)
	}

	// ShowBalanceStateOutputDTO is an example of output DTO at usecase level
	// actually, could not be needed!
	ShowBalanceStateOutputDTO struct{}

	ShowBalanceState struct {
		Repo ShowBalanceStateRepo
	}
)

func NewShowBalanceState(repo ShowBalanceStateRepo) *ShowBalanceState {
	return &ShowBalanceState{
		Repo: repo,
	}
}

func (s *ShowBalanceState) Execute(user *sharedkernel.User) (*ShowBalanceStateOutputDTO, error) {
	_, err := s.Repo.GetAccountByID(user.ID())
	if err != nil {
		return nil, err // nolint:wrapcheck // ok
	}

	core.NewAccount(user.ID())

	return &ShowBalanceStateOutputDTO{}, nil
}
