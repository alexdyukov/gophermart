package usecase

import (
	"context"
	"errors"

	"github.com/alexdyukov/gophermart/internal/gophermart/auth/domain/core"
	"golang.org/x/crypto/bcrypt"
)

type (
	LoginUserRepository interface {
		FindUserByLogin(context.Context, string) (*core.Credentials, error)
	}

	LoginUserInputPort interface {
		Execute(context.Context, UserInputDTO) (string, error)
	}

	LoginUser struct {
		Repo         LoginUserRepository
		TokenGateway AuthTokenIssuerGateway
	}
)

var ErrBadCredentials = errors.New("login password bad pair")

func NewLoginUser(repo LoginUserRepository, tg AuthTokenIssuerGateway) *LoginUser {
	return &LoginUser{
		Repo:         repo,
		TokenGateway: tg,
	}
}

func (login *LoginUser) Execute(ctx context.Context, dto UserInputDTO) (string, error) {
	credentials, err := login.Repo.FindUserByLogin(ctx, dto.Login)
	if err != nil {
		return "", err // nolint:wrapcheck // ok
	}

	err = bcrypt.CompareHashAndPassword([]byte(credentials.HashedPassword), []byte(dto.Password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return "", ErrBadCredentials
		}

		return "", err //nolint:wrapcheck // ok
	}

	jwt, err := login.TokenGateway.IssueWithLoginAndID(credentials.Login, credentials.UID)
	if err != nil {
		return "", err // nolint:wrapcheck // ok
	}

	return jwt, nil
}
