package usecase

import (
	"github.com/alexdyukov/gophermart/internal/accrual/domain/core"
)

type RegisterMechanicInputPort interface {
	Execute(string, RegisterMechanicInputDTO) error
}

type RegisterMechanicRepository interface {
	SaveMechanic(core.RewardMechanic) error
}

type RegisterMechanicInputDTO struct {
	/* mechanic fields */
	pattern string
	reward  int
	typ     string
}
type RegisterMechanicOutputDTO struct{}

type RegisterMechanic struct {
	Repo RegisterMechanicRepository
}

func NewRegisterMechanic(repo RegisterMechanicRepository) *RegisterMechanic {
	return &RegisterMechanic{
		Repo: repo,
	}
}

func (r *RegisterMechanic) Execute(user string, dto RegisterMechanicInputDTO) error {
	// map DTO into Entity
	mechanic := core.NewRewardMechanic(dto.pattern, dto.reward, dto.typ)
	_ = r.Repo.SaveMechanic(mechanic)
	return nil
}
