package memory

// Store ...
type GophermartStore struct {
	uRep *UsrRepository
	oRep *OrdRepository
	wRep *WithDrwRepository
}

func NewGophermartStore() *GophermartStore {

	return &GophermartStore{
		uRep: &UsrRepository{users: make(map[int]*UsersModel)},
		oRep: &OrdRepository{orders: make(map[int]*OrdersModel)},
		wRep: &WithDrwRepository{withdrawals: make(map[int]*WithdrawalsModel)},
	}

}
