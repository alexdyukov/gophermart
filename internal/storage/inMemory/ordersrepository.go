package inMemory

import (
	"github.com/alexdyukov/gophermart/internal/storage"
)

type OrdRepository struct {
	orders map[int]*storage.OrdersModel
}

func (u OrdRepository) Set(s *storage.OrdersModel, su *storage.UsersModel) error {

	u.orders[s.Number] = s
	u.orders[s.Number].UsersModel = *su

	return nil
}

func (u OrdRepository) Get(user *storage.UsersModel) ([]*storage.OrdersModel, error) {

	um := make([]*storage.OrdersModel, 0)

	for _, ord := range u.orders {
		if ord.UsersModel == *user {
			um = append(um, ord)

		}
	}

	return um, nil
}

func (u OrdRepository) GetAllSums(user *storage.UsersModel) (int, error) {
	sum := 0
	for _, o := range u.orders {
		if o.UsersModel.ID == user.ID {
			sum += o.Sum
		}

	}
	return sum, nil
}
