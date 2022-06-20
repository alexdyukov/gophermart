package inMemory

import (
	"github.com/alexdyukov/gophermart/internal/storage"
)

type WithDrwRepository struct {
	withdrawals map[int]*storage.WithdrawalsModel
}

func (u WithDrwRepository) Set(s *storage.WithdrawalsModel, su *storage.UsersModel) error {

	u.withdrawals[s.Number] = s
	u.withdrawals[s.Number].UsersModel = *su

	return nil
}

func (u WithDrwRepository) Get(user *storage.UsersModel) ([]*storage.WithdrawalsModel, error) {

	um := make([]*storage.WithdrawalsModel, 0)

	for _, wds := range u.withdrawals {
		if wds.UsersModel == *user {
			um = append(um, wds)

		}
	}

	return um, nil
}
func (u WithDrwRepository) GetAllSums(user *storage.UsersModel) (int, error) {
	sum := 0
	for _, o := range u.withdrawals {
		if o.UsersModel.ID == user.ID {
			sum += o.Sum
		}

	}
	return sum, nil
}
