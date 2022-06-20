package storage

type Storage interface {
	Users() UsersRepository
	Orders() OrdersRepository
	Withdraws() WithdrawalsRepository
}
