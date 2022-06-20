package storage

type UsersRepository interface {
	Set(*UsersModel) error           // Create
	Get(string) (*UsersModel, error) // find
}

type OrdersRepository interface {
	Set(*OrdersModel, *UsersModel) error
	Get(*UsersModel) ([]*OrdersModel, error) // Получение списка загруженных номеров заказов для пользователя
	GetAllSums(*UsersModel) (int, error)
}

type WithdrawalsRepository interface {
	Set(*WithdrawalsModel, *UsersModel) error
	Get(*UsersModel) ([]*WithdrawalsModel, error) // Получение информации о выводе средств для пользователя
	GetAllSums(*UsersModel) (int, error)
}
