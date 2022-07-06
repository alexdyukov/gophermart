package usecase

import (
	"context"
	"time"

	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

type (
	ListUserWithdrawalsRepository interface {
		FindAccountByID(context.Context, string) (core.Account, error)
	}

	ListUserWithdrawalsInputPort interface {
		Execute(context.Context, *sharedkernel.User) ([]ListUserWithdrawalsOutputDTO, error)
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

func (list *ListUserWithdrawals) Execute(ctx context.Context, user *sharedkernel.User) (
	[]ListUserWithdrawalsOutputDTO, error,
) { // nolint:whitespace // conflict with gofumpt
	_, err := list.Repo.FindAccountByID(ctx, user.ID())
	if err != nil {
		return nil, err // nolint:wrapcheck // ok
	}

	// map entity to output

	return []ListUserWithdrawalsOutputDTO{{}}, nil
}
