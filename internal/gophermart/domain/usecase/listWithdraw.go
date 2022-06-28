package usecase

import (
	"context"
	"time"

	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
)

type ListWithdrawalsRepository interface {
	GetAccountByID(context.Context, string) (core.Account, error)
}

type ListWithdrawalsInputPort interface {
	Execute(context.Context, string) (ListWithdrawalsOutputDTO, error)
}

type ListWithdrawalsOutputDTO struct {
	Order       int
	Sum         int
	ProcessedAt time.Time
}

type ListWithdrawals struct {
	Repo ListWithdrawalsRepository
}

func NewListWithdrawals(repo ListWithdrawalsRepository) *ListWithdrawals {
	return &ListWithdrawals{
		Repo: repo,
	}
}

func (l *ListWithdrawals) Execute(ctx context.Context, id string) (ListWithdrawalsOutputDTO, error) {
	_, _ = l.Repo.GetAccountByID(ctx, id)
	// prepare DTO response
	return ListWithdrawalsOutputDTO{}, nil
}
