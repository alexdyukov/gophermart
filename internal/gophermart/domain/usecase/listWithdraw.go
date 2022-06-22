package usecase

import (
	"time"

	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
)

type ListWithdrawalsRepository interface {
	GetAccountByID(string) (core.Account, error)
}

type ListWithdrawalsInputPort interface {
	Execute(string) (ListWithdrawalsOutputDTO, error)
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

func (l *ListWithdrawals) Execute(id string) (ListWithdrawalsOutputDTO, error) {
	_, _ = l.Repo.GetAccountByID(id)
	// prepare DTO response
	return ListWithdrawalsOutputDTO{}, nil
}
