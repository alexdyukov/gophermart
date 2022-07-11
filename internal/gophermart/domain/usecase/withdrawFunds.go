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
		Order string             `json:"order"`
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
	const (
		base    = 10
		bitSize = 64
	)

	if !sharedkernel.ValidLuhn(dto.Order) {
		return sharedkernel.ErrIncorrectOrderNumber
	}

	orderNumberInt, err := strconv.ParseInt(dto.Order, base, bitSize)
	if err != nil {
		return sharedkernel.ErrIncorrectOrderNumber
	}

	account, err := wuf.Repo.FindAccountByID(ctx, user.ID())
	if err != nil {
		return err
	}

	// do work with account, check if there is such number order in withdrarwals
	sliceAccountWithdrawals := core.GetSliceAccountWithdrawals(&account)
	for _, withdraw := range *sliceAccountWithdrawals {
		if withdraw.OrderNumber == orderNumberInt {
			return sharedkernel.ErrIncorrectOrderNumber
		}
	}

	err = account.WithdrawPoints(orderNumberInt, dto.Sum)
	if err != nil {
		return sharedkernel.ErrInsufficientFunds
	}

	err = wuf.Repo.SaveAccount(ctx, &account)
	if err != nil {
		return err
	}

	return nil
}
