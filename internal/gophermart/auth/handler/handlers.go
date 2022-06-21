package handler

import (
	"net/http"

	"github.com/alexdyukov/gophermart/internal/accrual/domain/usecase"
)

// PostAuth POST /api/user/login — аутентификация пользователя;
func PostAuth() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		//err := uc.Execute()
		//if err != nil {
		//	// todo: log error
		//	// todo: prepare response
		//}
		//200 — пользователь успешно аутентифицирован;
		//400 — неверный формат запроса;
		//401 — неверная пара логин/пароль;
		//500 — внутренняя ошибка сервера.
		writer.WriteHeader(200)
	}
}

// PostRegister POST /api/user/register — регистрация пользователя;
func PostRegister(uc usecase.CalculateLoyaltyPointsInputPort) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		//err := uc.Execute(dto)
		//if err != nil {
		//	// todo: log error
		//	// todo: prepare response
		//}
		//200 — пользователь успешно зарегистрирован и аутентифицирован;
		//400 — неверный формат запроса;
		//409 — логин уже занят;
		//500 — внутренняя ошибка сервера.
		writer.WriteHeader(200)
	}
}

func RefreshToken() {}
