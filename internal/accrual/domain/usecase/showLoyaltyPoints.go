package usecase

import (
	"context"
	"errors"
	"github.com/alexdyukov/gophermart/internal/accrual/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
	"strconv"
)

type ShowLoyaltyPointsRepository interface {
	GetLoyaltyPointsByOrderNumber(int, context.Context) (core.Answer, error)
}

type ShowLoyaltyPointsInputPort interface {
	Execute(int, context.Context) (core.Answer, error)
}

type ShowLoyaltyPoints struct {
	Repo ShowLoyaltyPointsRepository
}

func NewShowLoyaltyPoints(repo ShowLoyaltyPointsRepository) *ShowLoyaltyPoints {
	return &ShowLoyaltyPoints{
		Repo: repo,
	}
}

func (s *ShowLoyaltyPoints) Execute(number int, ctx context.Context) (core.Answer, error) {
	select {
	case <-ctx.Done():
		return core.Answer{}, nil
	default:
		if sharedkernel.ValidLuhn(strconv.Itoa(number)) {
			str, err := s.Repo.GetLoyaltyPointsByOrderNumber(number, ctx)
			return str, err
		}
		return core.Answer{}, errors.New("Номер заказа не валиден. не удовлетворяет алгоритму Луна")
	}

}
