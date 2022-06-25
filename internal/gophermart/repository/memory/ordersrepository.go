package memory

type OrdRepository struct {
	orders map[int]*OrdersModel
}

func (u OrdRepository) Set(s *OrdersModel, su *UsersModel) error {

	u.orders[s.Number] = s
	u.orders[s.Number].UsersModel = *su

	return nil
}

func (u OrdRepository) Get(user *UsersModel) ([]*OrdersModel, error) {

	um := make([]*OrdersModel, 0)

	for _, ord := range u.orders {
		if ord.UsersModel == *user {
			um = append(um, ord)

		}
	}

	return um, nil
}

func (u OrdRepository) GetAllSums(user *UsersModel) (int, error) {
	sum := 0
	for _, o := range u.orders {
		if o.UsersModel.ID == user.ID {
			sum += o.Sum
		}

	}
	return sum, nil
}
