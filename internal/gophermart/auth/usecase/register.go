package usecase

import "github.com/alexdyukov/gophermart/internal/sharedkernel"

type RegisterUserRepository interface {
	SaveUser(sharedkernel.User) error
}

type RegisterUser struct {
	Repo RegisterUserRepository
}

func (r *RegisterUser) Execute() error {
	return nil
}
