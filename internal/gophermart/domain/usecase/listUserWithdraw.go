package usecase

import (
	"context"
	"strconv"
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
		ProcessedAt string             `json:"processed_at"` // nolint:tagliatelle // external requirements
		Order       string             `json:"order"`
		Sum         sharedkernel.Money `json:"sum"`
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
	acc, err := list.Repo.FindAccountByID(ctx, user.ID())
	if err != nil {
		return nil, err // nolint:wrapcheck // ok
	}

	sliceAccountWithdrawals := core.GetSliceAccountWithdrawals(&acc)
	slwoDTO := make([]ListUserWithdrawalsOutputDTO, 0, len(*sliceAccountWithdrawals))

	for _, withdraw := range *sliceAccountWithdrawals {
		slwoDTO = append(slwoDTO,
			ListUserWithdrawalsOutputDTO{
				Order:       strconv.FormatInt(withdraw.OrderNumber, 10),
				Sum:         withdraw.Amount,
				ProcessedAt: withdraw.OperationTime.Format(time.RFC3339),
			})
	}

	return slwoDTO, nil
}
