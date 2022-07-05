package usecase

import (
	"time"

	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

type (
	ListUserWithdrawalsRepository interface {
		GetAccountByID(string) (core.Account, error)
	}

	ListUserWithdrawalsInputPort interface {
		Execute(user *sharedkernel.User) ([]ListUserWithdrawalsOutputDTO, error)
	}

	ListUserWithdrawalsOutputDTO struct {
		ProcessedAt time.Time
		Order       int
		Sum         int
	}

	ListUserWithdrawals struct {
		Repo ListUserWithdrawalsRepository
	}
)

func NewListUserWithdrawals(repo ListUserWithdrawalsRepository) *ListUserWithdrawals {
	return &ListUserWithdrawals{
		Repo: repo,
	}
}

func (l *ListUserWithdrawals) Execute(user *sharedkernel.User) ([]ListUserWithdrawalsOutputDTO, error) {
	_, err := l.Repo.GetAccountByID(user.ID())
	if err != nil {
		return nil, err // nolint:wrapcheck // ok
	}

	return []ListUserWithdrawalsOutputDTO{
		{},
	}, nil
}
