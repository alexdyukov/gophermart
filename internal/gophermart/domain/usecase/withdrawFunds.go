package usecase

import (
	"context"
	"strconv"

	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

type (
	WithdrawUserFundsRepository interface {
		FindAccountByID(context.Context, string) (core.Account, error)
		SaveAccount(context.Context, *core.Account) error
	}

	WithdrawFundsInputPort interface {
		Execute(context.Context, *sharedkernel.User, WithdrawUserFundsInputDTO) error
	}

	// WithdrawUserFundsInputDTO Example of DTO with json at usecase level, which not quite correct.
	WithdrawUserFundsInputDTO struct {
		Order int64              `json:"order"`
		Sum   sharedkernel.Money `json:"sum"`
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

func (wuf *WithdrawUserFunds) Execute(
	ctx context.Context,
	user *sharedkernel.User,
	dto WithdrawUserFundsInputDTO,
) error {
	//
	const base = 10

	if !sharedkernel.ValidLuhn(strconv.FormatInt(dto.Order, base)) {
		return sharedkernel.ErrIncorrectOrderNumber
	}

	account, err := wuf.Repo.FindAccountByID(ctx, user.ID())
	if err != nil {
		return err
	}

	// do work with account
	err = account.WithdrawPoints(dto.Order, dto.Sum)
	if err != nil {
		return sharedkernel.ErrInsufficientFunds
	}

	err = wuf.Repo.SaveAccount(ctx, &account)
	if err != nil {
		return err
	}

	return nil
}
