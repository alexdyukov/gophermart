package usecase

import (
	"context"

	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

type (
	ShowUserBalanceRepository interface {
		FindAccountByID(context.Context, string) (core.Account, error)
	}

	ShowUserBalancePrimaryPort interface {
		Execute(context.Context, *sharedkernel.User) (*ShowUserBalanceOutputDTO, error)
	}

	ShowUserBalanceOutputDTO struct {
		Current   sharedkernel.Money `json:"current"`
		Withdrawn sharedkernel.Money `json:"withdrawn"`
	}

	ShowUserBalance struct {
		Repo ShowUserBalanceRepository
	}
)

func NewShowUserBalance(repo ShowUserBalanceRepository) *ShowUserBalance {
	return &ShowUserBalance{
		Repo: repo,
	}
}

func (s *ShowUserBalance) Execute(ctx context.Context, user *sharedkernel.User) (*ShowUserBalanceOutputDTO, error) {
	userAccount, err := s.Repo.FindAccountByID(ctx, user.ID())
	if err != nil {
		return nil, err // nolint:wrapcheck // ok
	}

	output := ShowUserBalanceOutputDTO{
		Current:   userAccount.CurrentBalance(),
		Withdrawn: userAccount.WithdrawalsSum(),
	}

	return &output, nil
}
