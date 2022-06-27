package usecase

import (
	"context"
	"errors"
	"github.com/alexdyukov/gophermart/internal/accrual/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
	"strconv"
)

type ShowLoyaltyPointsRepository interface {
	GetLoyaltyPointsByOrderNumber(context.Context, int) (core.Answer, error)
}

type ShowLoyaltyPointsInputPort interface {
	Execute(context.Context, int) (core.Answer, error)
}

type ShowLoyaltyPoints struct {
	Repo ShowLoyaltyPointsRepository
}

func NewShowLoyaltyPoints(repo ShowLoyaltyPointsRepository) *ShowLoyaltyPoints {
	return &ShowLoyaltyPoints{
		Repo: repo,
	}
}

func (s *ShowLoyaltyPoints) Execute(ctx context.Context, number int) (core.Answer, error) {

	if sharedkernel.ValidLuhn(strconv.Itoa(number)) {
		str, err := s.Repo.GetLoyaltyPointsByOrderNumber(ctx, number)
		return str, err
	}
	return core.Answer{}, errors.New("Номер заказа не валиден. не удовлетворяет алгоритму Луна")

}
