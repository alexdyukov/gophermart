package usecase

import (
	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

type (
	WithdrawUserFundsRepository interface {
		GetAccountByID(string) (core.Account, error)
		SaveAccount(core.Account) error
	}

	WithdrawFundsInputPort interface {
		Execute(*sharedkernel.User, WithdrawUserFundsInputDTO) error
	}

	// WithdrawUserFundsInputDTO Example of DTO with json at usecase level, which not quite correct.
	WithdrawUserFundsInputDTO struct {
		Order int `json:"order"`
		Sum   int `json:"sum"`
	}

	WithdrawUserFunds struct {
		Repo WithdrawUserFundsRepository
	}
)

func NewWithdrawUserFunds(repo WithdrawUserFundsRepository) *WithdrawUserFunds {
	return &WithdrawUserFunds{
		Repo: repo,
	}
}

func (w *WithdrawUserFunds) Execute(user *sharedkernel.User, _ WithdrawUserFundsInputDTO) error {
	_, err := w.Repo.GetAccountByID(user.ID())
	if err != nil {
		return err // nolint:wrapcheck // ok
	}

	// do work with account
	// _ = account.WithdrawPoints(dto.Order, dto.Sum)
	// _ = w.Repository.SaveAccount(account)

	return nil
}
