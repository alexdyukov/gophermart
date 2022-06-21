package usecase

import "github.com/alexdyukov/gophermart/internal/sharedkernel"

type LoginUserRepository interface {
	FindUserByID(string) (sharedkernel.User, error)
}

type LoginUser struct {
	Repo LoginUserRepository
}

func (r *LoginUser) Execute(id string) error {
	return nil
}
