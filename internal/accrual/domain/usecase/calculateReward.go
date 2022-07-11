package usecase

import (
	"context"
	"errors"
	"regexp"

	"github.com/alexdyukov/gophermart/internal/accrual/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

type (
	CalculationRewardRepository interface {
		GetOrderByNumberWithGoods(context.Context, int64) (*core.OrderReceipt, error)
		FindAllRewardMechanicsByTokens(context.Context, ...string) (map[string]core.Reward, error)
		UpdateReceiptOrderState(context.Context, *core.OrderReceipt) error
	}

	CalculateRewardPrimaryPort interface {
		Execute(context.Context, *core.OrderReceipt) error
	}

	CalculateReward struct {
		Repo        CalculationRewardRepository
		brandFinder *regexp.Regexp
	}
)

var ErrNoRewards = errors.New("no rewards found by criteria")

func NewCalculateReward(repo CalculationRewardRepository) *CalculateReward {
	finder := regexp.MustCompile("[A-Z0-9a-z-]+")

	return &CalculateReward{
		Repo:        repo,
		brandFinder: finder,
	}
}

func (calc *CalculateReward) Execute(ctx context.Context, orderReceipt *core.OrderReceipt) error {
	orderReceipt.Status = sharedkernel.PROCESSING

	err := calc.Repo.UpdateReceiptOrderState(ctx, orderReceipt)
	if err != nil {
		return err
	}

	tokens := make([]string, 0, len(orderReceipt.Goods))

	for k, v := range orderReceipt.Goods {
		token := string(calc.brandFinder.Find([]byte(v.Description)))
		tokens = append(tokens, token)
		orderReceipt.Goods[k].Match = token
	}

	rewards, err := calc.Repo.FindAllRewardMechanicsByTokens(ctx, tokens...)
	if err != nil {
		if errors.Is(err, ErrNoRewards) {
			orderReceipt.Status = sharedkernel.PROCESSED

			return nil
		}

		return err
	}

	orderReceipt.CalculateRewardPoints(rewards)

	err = calc.Repo.UpdateReceiptOrderState(ctx, orderReceipt)
	if err != nil {
		return err
	}

	return nil
}
