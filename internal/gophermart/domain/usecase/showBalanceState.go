package usecase

import (
	"context"
	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
)

type ShowBalanceStateRepo interface {
	GetAccountByID(context.Context, string) (core.Account, error)
}

type ShowBalanceStateInputPort interface {
	Execute(context.Context, string) (ShowBalanceStateOutputDTO, error)
}

type ShowBalanceStateOutputDTO struct {
	Current   float32 `json:"current"`
	Withdrawn float32 `json:"withdrawn"`
}

type ShowBalanceState struct {
	Repo ShowBalanceStateRepo
}

func NewShowBalanceState(repo ShowBalanceStateRepo) *ShowBalanceState {
	return &ShowBalanceState{
		Repo: repo,
	}
}

func (s *ShowBalanceState) Execute(ctx context.Context, id string) (ShowBalanceStateOutputDTO, error) {
	// checks..
	_, err := s.Repo.GetAccountByID(ctx, id)
	if err != nil {
		// process error
	}
	// construct output
	return ShowBalanceStateOutputDTO{ /* fill with data */ }, nil
}
