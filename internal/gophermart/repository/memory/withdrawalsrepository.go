package memory

type WithDrwRepository struct {
	withdrawals map[int]*WithdrawalsModel
}

func (u WithDrwRepository) Set(s *WithdrawalsModel, su *UsersModel) error {

	u.withdrawals[s.Number] = s
	u.withdrawals[s.Number].UsersModel = *su

	return nil
}

func (u WithDrwRepository) Get(user *UsersModel) ([]*WithdrawalsModel, error) {

	um := make([]*WithdrawalsModel, 0)

	for _, wds := range u.withdrawals {
		if wds.UsersModel == *user {
			um = append(um, wds)

		}
	}

	return um, nil
}
func (u WithDrwRepository) GetAllSums(user *UsersModel) (int, error) {
	sum := 0
	for _, o := range u.withdrawals {
		if o.UsersModel.ID == user.ID {
			sum += o.Sum
		}

	}
	return sum, nil
}
