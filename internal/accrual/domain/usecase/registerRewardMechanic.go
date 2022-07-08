package usecase

import (
	"context"
	"errors"

	"github.com/alexdyukov/gophermart/internal/accrual/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

type (
	RegisterRewardMechanicPrimaryPort interface {
		Execute(context.Context, *RegisterRewardMechanicInputDTO) error
	}

	RegisterRewardMechanicRepository interface {
		SaveRewardMechanic(context.Context, *core.Reward) error
	}

	RegisterRewardMechanicInputDTO struct {
		Match      string             `json:"match"`
		RewardType string             `json:"reward_type"` // nolint:tagliatelle // external requirements
		Reward     sharedkernel.Money `json:"reward"`
	}

	RegisterRewardMechanic struct {
		Repo RegisterRewardMechanicRepository
	}
)

var (
	ErrEmptyInput          = errors.New("input cant be empty")
	ErrZeroInput           = errors.New("input cant be zero or less")
	ErrRewardAlreadyExists = errors.New("reward already exists")
)

func NewRegisterRewardMechanic(repo RegisterRewardMechanicRepository) *RegisterRewardMechanic {
	return &RegisterRewardMechanic{
		Repo: repo,
	}
}

func (reg *RegisterRewardMechanic) Execute(ctx context.Context, dto *RegisterRewardMechanicInputDTO) error {
	err := reg.validateInput(dto)
	if err != nil {
		return err
	}

	reward, err := core.NewReward(dto.Match, dto.Reward, dto.RewardType)
	if err != nil {
		return err
	}

	err = reg.Repo.SaveRewardMechanic(ctx, reward)
	if err != nil {
		return err //nolint:wrapcheck // ok
	}

	return nil
}

func (reg *RegisterRewardMechanic) validateInput(dto *RegisterRewardMechanicInputDTO) error {
	if len(dto.Match) < 1 {
		return ErrEmptyInput
	}

	if dto.Reward < 1 {
		return ErrZeroInput
	}

	return nil
}
