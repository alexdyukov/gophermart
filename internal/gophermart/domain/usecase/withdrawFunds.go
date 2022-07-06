package usecase

import (
	"context"
	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
)

// Делала BeOl только добавляла контекст везде
// POST /api/user/balance/withdraw

type WithdrawFundsRepository interface {
	GetAccountByID(context.Context, string) (core.Account, error)
	SaveAccount(context.Context, core.Account) error
}

type WithdrawFundsInputPort interface {
	Execute(context.Context, string, WithdrawFundsInputDTO) error
}

// WithdrawFundsInputDTO Example of DTO with json at usecase level not quite correct
type WithdrawFundsInputDTO struct {
	Order string  `json:"order"`
	Sum   float32 `json:"sum"`
}

type WithdrawFunds struct {
	Repo WithdrawFundsRepository
}

func NewWithdrawFunds(repo WithdrawFundsRepository) *WithdrawFunds {
	return &WithdrawFunds{
		Repo: repo,
	}
}

func (w *WithdrawFunds) Execute(ctx context.Context, id string, dto WithdrawFundsInputDTO) error {

	account, _ := w.Repo.GetAccountByID(ctx, id)
	// do work with account
	_ = account.WithdrawPoints(dto.Order, dto.Sum)
	_ = w.Repo.SaveAccount(ctx, account)

	return nil
}
