package usecase

import (
	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

type (
	ShowBalanceStateRepo interface {
		GetAccountByID(string) (core.Account, error)
	}

	ShowUserBalanceInputPort interface {
		Execute(user *sharedkernel.User) (*ShowUserBalanceOutputDTO, error)
	}

	// ShowUserBalanceOutputDTO is an example of output DTO at usecase level
	// actually, could not be needed!
	ShowUserBalanceOutputDTO struct{}

	ShowUserBalance struct {
		Repo ShowBalanceStateRepo
	}
)

func NewShowBalanceState(repo ShowBalanceStateRepo) *ShowUserBalance {
	return &ShowUserBalance{
		Repo: repo,
	}
}

func (s *ShowUserBalance) Execute(user *sharedkernel.User) (*ShowUserBalanceOutputDTO, error) {
	_, err := s.Repo.GetAccountByID(user.ID())
	if err != nil {
		return nil, err // nolint:wrapcheck // ok
	}

	core.NewAccount(user.ID())

	return &ShowUserBalanceOutputDTO{}, nil
}
