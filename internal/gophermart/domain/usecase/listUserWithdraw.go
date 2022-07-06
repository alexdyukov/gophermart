package usecase

import (
	"time"

	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

type (
	ListUserWithdrawalsRepository interface {
		FindAccountByID(string) (core.Account, error)
	}

	ListUserWithdrawalsInputPort interface {
		Execute(user *sharedkernel.User) ([]ListUserWithdrawalsOutputDTO, error)
	}

	ListUserWithdrawalsOutputDTO struct {
		ProcessedAt time.Time `json:"processed_at"` // nolint:tagliatelle // external requirements
		Order       string    `json:"order"`
		Sum         int       `json:"sum"`
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
	_, err := l.Repo.FindAccountByID(user.ID())
	if err != nil {
		return nil, err // nolint:wrapcheck // ok
	}

	// map entity to output

	return []ListUserWithdrawalsOutputDTO{{}}, nil
}
