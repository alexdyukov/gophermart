package usecase

import (
	"context"
)

// Делала BeOl
type ShowBalanceStateRepo interface {
	GetBalance(context.Context, string) (float32, float32, error)
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
	currentAll, withdrawals, err := s.Repo.GetBalance(ctx, id)

	if err != nil {
		// process error
	}
	// construct output
	return ShowBalanceStateOutputDTO{Current: currentAll - withdrawals, Withdrawn: withdrawals}, nil
}
