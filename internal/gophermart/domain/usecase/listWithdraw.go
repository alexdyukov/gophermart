package usecase

import (
	"time"

	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

type (
	ListWithdrawalsRepository interface {
		GetAccountByID(string) (core.Account, error)
	}

	ListWithdrawalsInputPort interface {
		Execute(user *sharedkernel.User) (*ListWithdrawalsOutputDTO, error)
	}

	ListWithdrawalsOutputDTO struct {
		ProcessedAt time.Time
		Order       int
		Sum         int
	}

	ListWithdrawals struct {
		Repo ListWithdrawalsRepository
	}
)

func NewListWithdrawals(repo ListWithdrawalsRepository) *ListWithdrawals {
	return &ListWithdrawals{
		Repo: repo,
	}
}

func (l *ListWithdrawals) Execute(user *sharedkernel.User) (*ListWithdrawalsOutputDTO, error) {
	_, err := l.Repo.GetAccountByID(user.ID())
	if err != nil {
		return nil, err // nolint:wrapcheck // ok
	}

	return &ListWithdrawalsOutputDTO{}, nil // nolint:exhaustivestruct // ok
}
