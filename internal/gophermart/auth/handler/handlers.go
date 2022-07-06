package handler

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/alexdyukov/gophermart/internal/gophermart/auth/domain/usecase"
)

// LoginPostHandler POST /api/user/login — аутентификация пользователя.
// 200 — пользователь успешно аутентифицирован;
// 400 — неверный формат запроса;
// 401 — неверная пара логин/пароль;
// 500 — внутренняя ошибка сервера.
func LoginPostHandler(loginUsecase usecase.LoginUserInputPort) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		bytes, err := io.ReadAll(request.Body)
		if err != nil {
			log.Printf("error while reading request.")
			writer.WriteHeader(http.StatusBadRequest)

			return
		}

		inputData := usecase.UserInputDTO{} // nolint:exhaustivestruct // ok.

		err = json.Unmarshal(bytes, &inputData)
		if err != nil {
			log.Printf("error while unmarshalling json")
			writer.WriteHeader(http.StatusBadRequest)

			return
		}

		jwtString, err := loginUsecase.Execute(request.Context(), inputData)
		if err != nil {
			if errors.Is(err, usecase.ErrBadCredentials) {
				log.Println(err)
				writer.WriteHeader(http.StatusUnauthorized)

				return
			}

			log.Println(err)

			return
		}

		// Header or Cookie
		// writer.Header().Set("Authorization", jwtString)

		cookie := http.Cookie{ // nolint:exhaustivestruct // ok
			Name:   "auth",
			Value:  jwtString,
			MaxAge: 3600 * 24, //nolint:gomnd // temporary
		}

		http.SetCookie(writer, &cookie)
		writer.WriteHeader(http.StatusOK)
	}
}

// RegisterPostHandler POST /api/user/register — регистрация пользователя.
// 200 — пользователь успешно зарегистрирован и аутентифицирован
// 400 — неверный формат запроса;
// 409 — логин уже занят;
// 500 — внутренняя ошибка сервера.
func RegisterPostHandler(registerUsecase usecase.RegisterUserPrimaryPort) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		bytes, err := io.ReadAll(request.Body)
		if err != nil {
			log.Printf("error while reading request.")
			writer.WriteHeader(http.StatusBadRequest)

			return
		}

		inputData := usecase.UserInputDTO{} // nolint:exhaustivestruct // ok

		err = json.Unmarshal(bytes, &inputData)
		if err != nil {
			log.Printf("error while unmarshalling json")
			writer.WriteHeader(http.StatusBadRequest)

			return
		}

		jwtString, err := registerUsecase.Execute(request.Context(), inputData)
		if err != nil {
			if errors.Is(err, usecase.ErrLoginAlreadyExist) {
				writer.WriteHeader(http.StatusConflict)

				return
			}

			log.Println(err.Error())
			writer.WriteHeader(http.StatusInternalServerError)

			return
		}

		c := http.Cookie{
			Name:   "auth",
			Value:  jwtString,
			MaxAge: 3600 * 24, // nolint:gomnd // temporary until config
		}
		http.SetCookie(writer, &c)
		writer.WriteHeader(http.StatusOK)
	}
}
