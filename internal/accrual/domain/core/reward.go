package core

import "github.com/alexdyukov/gophermart/internal/sharedkernel"

type Reward struct {
	id           string
	match        string
	rewardType   string
	rewardPoints sharedkernel.Money
}

func NewReward(match string, rewardPoints sharedkernel.Money, rewardType string) Reward {
	return Reward{
		id:           sharedkernel.NewUUID(),
		match:        match,
		rewardPoints: rewardPoints,
		rewardType:   rewardType,
	}
}
