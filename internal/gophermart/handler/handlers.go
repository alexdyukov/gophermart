package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/alexdyukov/gophermart/internal/gophermart/domain/usecase"
	"github.com/alexdyukov/gophermart/internal/gophermart/handler/middleware"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

// RegisterUserOrderPostHandler POST /api/user/orders — загрузка пользователем номера заказа для расчёта.
// 200 — номер заказа уже был загружен этим пользователем;
// 202 — новый номер заказа принят в обработку;
// 400 — неверный формат запроса;
// 401 — пользователь не аутентифицирован;
// 409 — номер заказа уже был загружен другим пользователем;
// 422 — неверный формат номера заказа;
// 500 — внутренняя ошибка сервера.
func RegisterUserOrderPostHandler(registerUserOrderUsecase usecase.RegisterUserOrderPrimaryPort) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		user, ok := request.Context().Value(middleware.User).(*sharedkernel.User)
		if !ok {
			writer.WriteHeader(http.StatusUnauthorized)

			return
		}

		bytes, err := io.ReadAll(request.Body)
		if err != nil {
			log.Printf("error while reading request.")
			writer.WriteHeader(http.StatusBadRequest)

			return
		}

		orderNumber, err := strconv.Atoi(string(bytes))
		if err != nil {
			log.Printf("error while reading request.")
			writer.WriteHeader(http.StatusBadRequest)

			return
		}

		err = registerUserOrderUsecase.Execute(request.Context(), orderNumber, user)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)

			return
		}

		writer.WriteHeader(http.StatusOK)
	}
}

// ListUserOrdersGetHandler GET /api/user/orders — получение списка загруженных пользователем номеров заказов,
// статусов их обработки и информации о начислениях
// 200 — успешная обработка запроса.
// 204 — нет данных для ответа.
// 401 — пользователь не авторизован.
// 500 — внутренняя ошибка сервера.
func ListUserOrdersGetHandler(listUserOrdersUsecase usecase.ListUserOrdersPrimaryPort) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		user, ok := request.Context().Value(middleware.User).(*sharedkernel.User)
		if !ok {
			writer.WriteHeader(http.StatusUnauthorized)

			return
		}

		list, err := listUserOrdersUsecase.Execute(request.Context(), user)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError) // 500 — внутренняя ошибка сервера

			return
		}

		if len(list) == 0 {
			writer.WriteHeader(http.StatusNoContent) // 204 — нет данных для ответа.

			return
		}

		strJSON, err := json.Marshal(list)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError) // 500 — внутренняя ошибка сервера

			return
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK) // 200 — успешная обработка запроса.
		_, err = writer.Write(strJSON)

		if err != nil {
			log.Println(err)

			return
		}
	}
}

// GetBalance GET /api/user/balance — получение текущего баланса счёта баллов лояльности пользователя
// 200 — успешная обработка запроса.
// 401 — пользователь не авторизован.
// 500 — внутренняя ошибка сервера.
func GetBalance(showBalanceUsecase usecase.ShowUserBalancePrimaryPort) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		user, ok := request.Context().Value(middleware.User).(*sharedkernel.User)
		if !ok {
			writer.WriteHeader(http.StatusUnauthorized)

			return
		}

		balance, err := showBalanceUsecase.Execute(request.Context(), user)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)

			return
		}

		strJSON, err := json.Marshal(balance)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError) // 500 — внутренняя ошибка сервера

			return
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)

		_, err = writer.Write(strJSON)

		if err != nil {
			log.Println(err)

			return
		}
	}
}

// PostWithdraw POST /api/user/balance/withdraw — запрос на списание баллов с накопительного счёта
// в счёт оплаты нового заказа
// 200 — успешная обработка запроса;
// 401 — пользователь не авторизован;
// 402 — на счету недостаточно средств;
// 422 — неверный номер заказа;
// 500 — внутренняя ошибка сервера.
func PostWithdraw(withdrawFundsUsecase usecase.WithdrawFundsInputPort) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		user, ok := request.Context().Value(middleware.User).(*sharedkernel.User)
		if !ok {
			writer.WriteHeader(http.StatusUnauthorized)

			return
		}

		dto := usecase.WithdrawUserFundsInputDTO{} // nolint:exhaustivestruct // ok,  exhaustive // ok.

		bytes, err := io.ReadAll(request.Body)
		if err != nil {
			log.Println(err)
		}

		err = json.Unmarshal(bytes, &dto)
		if err != nil {
			log.Println(err)
		}

		err = withdrawFundsUsecase.Execute(request.Context(), user, dto)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)

			return
		}

		writer.WriteHeader(http.StatusOK)
	}
}

// GetWithdrawals GET /api/user/balance/withdrawals — получение информации о выводе средств с накопительного счёта
// 200 — успешная обработка запроса;
// 204 — нет ни одного списания.
// 401 — пользователь не авторизован.
// 500 — внутренняя ошибка сервера.
func GetWithdrawals(listWithdrawalsUsecase usecase.ListUserWithdrawalsInputPort) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		user, ok := request.Context().Value(middleware.User).(*sharedkernel.User)
		if !ok {
			writer.WriteHeader(http.StatusUnauthorized)

			return
		}

		_, err := listWithdrawalsUsecase.Execute(request.Context(), user)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)

			return
		}

		writer.WriteHeader(http.StatusOK)
	}
}
