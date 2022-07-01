package middleware

import (
	"context"
	"net/http"

	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

type (
	TokenValidator interface {
		ValidateWithLoginAndID(string) (*sharedkernel.User, error)
	}

	ctxKey int
)

const (
	User ctxKey = iota
)

func Authentication(tokenValidator TokenValidator) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			jwtCookie, err := request.Cookie("auth")
			if err != nil {
				writer.WriteHeader(http.StatusUnauthorized)

				return
			}

			user, err := tokenValidator.ValidateWithLoginAndID(jwtCookie.Value)
			if err != nil {
				writer.WriteHeader(http.StatusUnauthorized)

				return
			}

			userCtx := context.WithValue(request.Context(), User, user)
			request = request.WithContext(userCtx)

			next.ServeHTTP(writer, request)
		})
	}
}
