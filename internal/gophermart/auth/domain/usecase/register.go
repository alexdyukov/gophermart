package usecase

import (
	"context"
	"errors"

	"github.com/alexdyukov/gophermart/internal/sharedkernel"
	"golang.org/x/crypto/bcrypt"
)

type (
	RegisterUserRepository interface {
		SaveUserIfNotExist(context.Context, *sharedkernel.User, string) error
	}

	RegisterUserPrimaryPort interface {
		Execute(context.Context, UserInputDTO) (string, error)
	}

	RegisterUser struct {
		Repo         RegisterUserRepository
		TokenGateway AuthTokenIssuerGateway
	}
)

var ErrLoginAlreadyExist = errors.New("login is already taken")

func NewRegisterUser(repo RegisterUserRepository, tg AuthTokenIssuerGateway) *RegisterUser {
	return &RegisterUser{
		Repo:         repo,
		TokenGateway: tg,
	}
}

func (reg *RegisterUser) Execute(ctx context.Context, dto UserInputDTO) (string, error) {
	user := sharedkernel.NewUser(dto.Login)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), 0)
	if err != nil {
		return "", err // nolint:wrapcheck //ok
	}

	err = reg.Repo.SaveUserIfNotExist(ctx, user, string(hashedPassword))
	if err != nil {
		return "", err // nolint:wrapcheck // ok
	}

	jwt, err := reg.TokenGateway.IssueWithLoginAndID(user.Login(), user.ID())
	if err != nil {
		return "", err // nolint:wrapcheck // ok
	}

	return jwt, nil
}
