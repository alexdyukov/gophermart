package usecase

import (
	"github.com/alexdyukov/gophermart/internal/accrual/domain/core"
)

type (
	RegisterMechanicInputPort interface {
		Execute(string, *RegisterMechanicInputDTO) error
	}

	RegisterMechanicRepository interface {
		SaveMechanic(*core.RewardMechanic) error
	}

	RegisterMechanicInputDTO struct {
		pattern string
		typ     string
		reward  int
	}

	RegisterMechanicOutputDTO struct{}

	RegisterMechanic struct {
		Repo RegisterMechanicRepository
	}
)

func NewRegisterMechanic(repo RegisterMechanicRepository) *RegisterMechanic {
	return &RegisterMechanic{
		Repo: repo,
	}
}

func (r *RegisterMechanic) Execute(_ string, dto *RegisterMechanicInputDTO) error {
	mechanic := core.NewRewardMechanic(dto.pattern, dto.reward, dto.typ)

	err := r.Repo.SaveMechanic(&mechanic)
	if err != nil {
		return err //nolint:wrapcheck // ok
	}

	return nil
}
