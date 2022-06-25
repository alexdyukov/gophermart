package usecase

import "github.com/alexdyukov/gophermart/internal/gophermart/domain/core"

type WithdrawFundsRepository interface {
	GetAccountByID(string) (core.Account, error)
	SaveAccount(core.Account) error
}

type WithdrawFundsInputPort interface {
	Execute(string, WithdrawFundsInputDTO) error
}

// WithdrawFundsInputDTO Example of DTO with json at usecase level not quite correct
type WithdrawFundsInputDTO struct {
	Order int `json:"order"`
	Sum   int `json:"sum"`
}

type WithdrawFunds struct {
	Repo WithdrawFundsRepository
}

func NewWithdrawFunds(repo WithdrawFundsRepository) *WithdrawFunds {
	return &WithdrawFunds{
		Repo: repo,
	}
}

func (w *WithdrawFunds) Execute(id string, dto WithdrawFundsInputDTO) error {

	account, _ := w.Repo.GetAccountByID(id)
	// do work with account
	_ = account.WithdrawPoints(dto.Order, dto.Sum)
	_ = w.Repo.SaveAccount(account)

	return nil
}
