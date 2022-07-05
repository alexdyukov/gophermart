package postgres

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/alexdyukov/gophermart/internal/gophermart/auth/domain/core"
	"github.com/alexdyukov/gophermart/internal/gophermart/auth/domain/usecase"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
	"github.com/jackc/pgconn"
)

type AuthStore struct {
	db *sql.DB
}

func NewAuthStore(db *sql.DB) (*AuthStore, error) {
	authStore := AuthStore{
		db: db,
	}

	err := authStore.createUserTableIfNotExist()
	if err != nil {
		return nil, err
	}

	return &authStore, nil
}

// SaveUserIfNotExist contains tx for future usage If it will be needed to add another SQL stmts
// checking queries etc. For using another saving mechanic.
func (a *AuthStore) SaveUserIfNotExist(ctx context.Context, user *sharedkernel.User, hashed string) error {
	transaction, err := a.db.Begin()
	if err != nil {
		return err // nolint:wrapcheck // ok
	}

	defer func() {
		err = transaction.Rollback()
		if err != nil {
			log.Println(err)
		}
	}()

	stmt, err := transaction.Prepare(`INSERT INTO public.user (uid, login, passwd) VALUES ($1,$2,$3)`)
	if err != nil {
		return err // nolint:wrapcheck // ok
	}

	defer func() {
		err = stmt.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	_, err = stmt.ExecContext(ctx, user.ID(), user.Login(), hashed)
	if err != nil {
		var pgError *pgconn.PgError
		if ok := errors.As(err, &pgError); ok {
			if pgError.Code == "23505" {
				return usecase.ErrLoginAlreadyExist
			}

			return err // nolint:wrapcheck // ok
		}

		return err // nolint:wrapcheck // ok
	}

	err = transaction.Commit()
	if err != nil {
		return err // nolint:wrapcheck // ok
	}

	return nil
}

func (a *AuthStore) FindUserByLogin(ctx context.Context, login string) (*core.Credentials, error) {
	stmt, err := a.db.PrepareContext(ctx, `SELECT uid, login, passwd FROM public.user WHERE login  = $1 LIMIT 1`)
	if err != nil {
		return nil, err // nolint:wrapcheck // ok
	}

	credentials := core.Credentials{} // nolint:exhaustivestruct // ok

	err = stmt.QueryRowContext(ctx, login).Scan(&credentials.UID, &credentials.Login, &credentials.HashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, usecase.ErrBadCredentials
		}

		return nil, err //nolint:wrapcheck  // ok
	}

	defer func() {
		err := stmt.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	return &credentials, nil
}

func (a *AuthStore) createUserTableIfNotExist() error {
	_, err := a.db.Exec(`CREATE TABLE IF NOT EXISTS public.user (
												uid TEXT NOT NULL,
												login TEXT NOT NULL,
												passwd TEXT NOT NULL,
												CONSTRAINT auth_pk_constraint PRIMARY KEY (uid),
												CONSTRAINT login_uniq_constraint UNIQUE (login));
												`)
	if err != nil {
		return err // nolint:wrapcheck // ok
	}

	return nil
}
