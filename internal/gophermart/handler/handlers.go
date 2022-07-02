package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/alexdyukov/gophermart/internal/gophermart/domain/usecase"
)

// PostOrder POST /api/user/orders — загрузка пользователем номера заказа для расчёта;
func PostOrder(uc usecase.LoadOrderNumberInputPort) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// prepare dto
		orderNum := 0
		err := uc.Execute(orderNum)
		if err != nil {
			// todo: log
			// todo: prepare response
		}
		//200 — номер заказа уже был загружен этим пользователем;
		//202 — новый номер заказа принят в обработку;
		//400 — неверный формат запроса;
		//401 — пользователь не аутентифицирован;
		//409 — номер заказа уже был загружен другим пользователем;
		//422 — неверный формат номера заказа;
		//500 — внутренняя ошибка сервера.
		writer.WriteHeader(200)
	}
}

// GetOrders GET /api/user/orders — получение списка загруженных пользователем номеров заказов, статусов их обработки и информации о начислениях;
func GetOrders(uc usecase.ListOrderNumsInputPort) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		user := "1" // тут мы должны будем получить пользователя после авторизации
		list, err := uc.Execute(request.Context(), user)
		if err != nil {
			//204 — нет данных для ответа.
			//401 — пользователь не авторизован.
			//500 — внутренняя ошибка сервера.
			// взависимости от полученной ошибки возвращаем тот или иной ответ, пока не ясно как эти ошибки получаем
			return
		}

		//200 — успешная обработка запроса.
		writer.WriteHeader(200)
		writer.Header().Set("Content-Type", "application/json")
		strJSON, err := json.Marshal(list)

		_, err = writer.Write(strJSON)

	}
}

// GetBalance GET /api/user/balance — получение текущего баланса счёта баллов лояльности пользователя;
func GetBalance(uc usecase.ShowBalanceStateInputPort) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// делаю эту процедуру
		user := "1" // тут мы должны будем получить пользователя после авторизации
		// наверное получаем  данные из request??

		balance, err := uc.Execute(request.Context(), user)
		if err != nil {
			//401 — пользователь не авторизован.
			//500 — внутренняя ошибка сервера.
			return
		}

		//200 — успешная обработка запроса.
		writer.WriteHeader(200)
		writer.Header().Set("Content-Type", "application/json")
		strJSON, err := json.Marshal(balance)
		_, err = writer.Write(strJSON)

	}
}

// PostWithdraw POST /api/user/balance/withdraw — запрос на списание баллов с накопительного счёта в счёт оплаты нового заказа;
func PostWithdraw(uc usecase.WithdrawFundsInputPort) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		user := "some user"
		dto := usecase.WithdrawFundsInputDTO{}

		// Unmarshal into DTO
		bytes, _ := io.ReadAll(request.Body)
		_ = json.Unmarshal(bytes, &dto)

		err := uc.Execute(request.Context(), user, dto)
		if err != nil {
			// todo: log
			// todo: prepare response
			return
		}
		//200 — успешная обработка запроса;
		//401 — пользователь не авторизован;
		//402 — на счету недостаточно средств;
		//422 — неверный номер заказа;
		//500 — внутренняя ошибка сервера.
		writer.WriteHeader(200)
	}
}

// GetWithdrawals GET /api/user/balance/withdrawals — получение информации о выводе средств с накопительного счёта пользователем.
func GetWithdrawals(uc usecase.ListWithdrawalsInputPort) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		user := "some user"
		_, err := uc.Execute(request.Context(), user)
		if err != nil {
			// todo: log
			// todo: prepare response
			return
		}
		//200 — успешная обработка запроса;
		//204 — нет ни одного списания.
		//401 — пользователь не авторизован.
		//500 — внутренняя ошибка сервера.
		writer.WriteHeader(200)
	}
}
