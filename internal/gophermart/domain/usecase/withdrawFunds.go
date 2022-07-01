package usecase

import (
	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

type (
	WithdrawFundsRepository interface {
		GetAccountByID(string) (core.Account, error)
		SaveAccount(core.Account) error
	}

	WithdrawFundsInputPort interface {
		Execute(*sharedkernel.User, WithdrawFundsInputDTO) error
	}

	// WithdrawFundsInputDTO Example of DTO with json at usecase level, which not quite correct.
	WithdrawFundsInputDTO struct {
		Order int `json:"order"`
		Sum   int `json:"sum"`
	}

	WithdrawFunds struct {
		Repo WithdrawFundsRepository
	}
)

func NewWithdrawFunds(repo WithdrawFundsRepository) *WithdrawFunds {
	return &WithdrawFunds{
		Repo: repo,
	}
}

func (w *WithdrawFunds) Execute(user *sharedkernel.User, _ WithdrawFundsInputDTO) error {
	_, err := w.Repo.GetAccountByID(user.ID())
	if err != nil {
		return err // nolint:wrapcheck // ok
	}

	// do work with account
	// _ = account.WithdrawPoints(dto.Order, dto.Sum)
	// _ = w.Repo.SaveAccount(account)

	return nil
}
