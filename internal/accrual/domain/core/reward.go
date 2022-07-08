package core

import (
	"errors"

	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

type Reward struct {
	match        string
	rewardType   string
	rewardPoints sharedkernel.Money
}

var ErrInvalidRewardType = errors.New("this reward type not supported")

func (r *Reward) Match() string {
	return r.match
}

func (r *Reward) RewardType() string {
	return r.rewardType
}

func (r *Reward) RewardPoints() sharedkernel.Money {
	return r.rewardPoints
}

func (r *Reward) isPercentage() bool {
	return r.RewardType() == "%"
}

func NewReward(match string, rewardPoints sharedkernel.Money, rewardType string) (*Reward, error) {
	err := checkRewardType(rewardType)
	if err != nil {
		return nil, err
	}

	reward := Reward{
		match:        match,
		rewardPoints: rewardPoints,
		rewardType:   rewardType,
	}

	return &reward, nil
}

func RestoreReward(match string, rewardPoints sharedkernel.Money, rewardType string) *Reward {
	reward := Reward{
		match:        match,
		rewardPoints: rewardPoints,
		rewardType:   rewardType,
	}

	return &reward
}

func checkRewardType(rewardType string) error {
	if (rewardType != "%") && (rewardType != "pt") {
		return ErrInvalidRewardType
	}

	return nil
}
