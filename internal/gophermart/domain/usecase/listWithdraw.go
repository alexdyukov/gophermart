package usecase

import (
	"context"
	"sort"
	"time"

	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
)

//Делала BeOl
//3- GET /api/user/balance/withdrawals
type ListWithdrawalsRepository interface {
	GetAccountByID(context.Context, string) (core.Account, error)
}

type ListWithdrawalsInputPort interface {
	Execute(context.Context, string) ([]ListWithdrawalsOutputDTO, error)
}

type ListWithdrawalsOutputDTO struct {
	Order          int       `json:"order"`
	Sum            int       `json:"sum"`
	ProcessedAt    time.Time `json:"-"`
	ProcessedAtStr string    `json:"processed_at"`
}

type ListWithdrawals struct {
	Repo ListWithdrawalsRepository
}

func NewListWithdrawals(repo ListWithdrawalsRepository) *ListWithdrawals {
	return &ListWithdrawals{
		Repo: repo,
	}
}

func (l *ListWithdrawals) Execute(ctx context.Context, id string) ([]ListWithdrawalsOutputDTO, error) {
	slwoDTO := make([]ListWithdrawalsOutputDTO, 0)

	acc, err := l.Repo.GetAccountByID(ctx, id)
	if err != nil {
		return slwoDTO, err
	}

	for _, withdraw := range acc.WithdrawHistory {
		slwoDTO = append(slwoDTO, ListWithdrawalsOutputDTO{Order: withdraw.OrderNumber, Sum: withdraw.Amount, ProcessedAt: withdraw.Time, ProcessedAtStr: withdraw.Time.Format(time.RFC3339)})
	}

	sort.SliceStable(slwoDTO, func(i, j int) bool {
		return slwoDTO[i].ProcessedAt.Before(slwoDTO[j].ProcessedAt)
	})

	return slwoDTO, nil

}
