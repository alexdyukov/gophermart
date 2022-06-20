package inMemory

import "github.com/alexdyukov/gophermart/internal/storage"

// Store ...
type Store struct {
	uRep *UsrRepository
	oRep *OrdRepository
	wRep *WithDrwRepository
}

func New() *Store {

	return &Store{}

}

func (s *Store) Users() *UsrRepository {
	if s.uRep != nil {
		return s.uRep
	}

	s.uRep = &UsrRepository{
		users: make(map[int]*storage.UsersModel),
	}

	return s.uRep
}

func (s *Store) Orders() *OrdRepository {
	if s.oRep != nil {
		return s.oRep
	}

	s.oRep = &OrdRepository{
		orders: make(map[int]*storage.OrdersModel),
	}

	return s.oRep
}

func (s *Store) Withdraws() *WithDrwRepository {
	if s.wRep != nil {
		return s.wRep
	}

	s.wRep = &WithDrwRepository{
		withdrawals: make(map[int]*storage.WithdrawalsModel),
	}

	return s.wRep
}
