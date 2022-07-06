package usecase

import (
	"github.com/alexdyukov/gophermart/internal/accrual/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

type (
	RegisterRewardMechanicPrimaryPort interface {
		Execute(string, *RegisterRewardMechanicInputDTO) error
	}

	RegisterRewardMechanicRepository interface {
		SaveRewardMechanic(*core.Reward) error
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

func NewRegisterRewardMechanic(repo RegisterRewardMechanicRepository) *RegisterRewardMechanic {
	return &RegisterRewardMechanic{
		Repo: repo,
	}
}

func (r *RegisterRewardMechanic) Execute(_ string, dto *RegisterRewardMechanicInputDTO) error {
	mechanic := core.NewReward(dto.Match, dto.Reward, dto.RewardType)

	err := r.Repo.SaveRewardMechanic(&mechanic)
	if err != nil {
		return err //nolint:wrapcheck // ok
	}

	return nil
}
