package handler

import (
	"encoding/json"
	"fmt"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
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

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(200)
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

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(200)
		strJSON, err := json.Marshal(balance)
		_, err = writer.Write(strJSON)

	}
}

// PostWithdraw POST /api/user/balance/withdraw — запрос на списание баллов с накопительного счёта в счёт оплаты нового заказа;
func PostWithdraw(uc usecase.WithdrawFundsInputPort) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("PostWithdraw: запустился хендлер  /api/user/balance/withdraw")
		//пример из тестов  66340157222 324.82
		user := "some user"
		dto := usecase.WithdrawFundsInputDTO{}

		// Unmarshal into DTO
		bytes, _ := io.ReadAll(request.Body)
		_ = json.Unmarshal(bytes, &dto)

		fmt.Println("PostWithdraw: получили запрос на списание, такие данные", dto)

		if sharedkernel.ValidLuhn(dto.Order) {

			err := uc.Execute(request.Context(), user, dto)
			if err != nil {
				fmt.Println("PostWithdraw: ошибка №1", err)
				writer.WriteHeader(500)
				return
			}
		} else {
			fmt.Println("PostWithdraw: неверный номер заказа - не отработал алгоритм луна")
			writer.WriteHeader(422)
		}
		//200 — успешная обработка запроса;
		//401 — пользователь не авторизован;
		//402 — на счету недостаточно средств;
		//422 — неверный номер заказа;
		//500 — внутренняя ошибка сервера.
		fmt.Println("PostWithdraw: все зашибись, отправляем статус 200")
		writer.WriteHeader(200)
	}
}

// GetWithdrawals GET /api/user/withdrawals — получение информации о выводе средств с накопительного счёта пользователем. // убрала balance
func GetWithdrawals(uc usecase.ListWithdrawalsInputPort) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		user := "some user"
		fmt.Println("GetWithdrawals: запустился хендлер  /api/user/withdrawals")
		wdrls, err := uc.Execute(request.Context(), user)

		if err != nil {
			fmt.Println("GetWithdrawals: ушли в ошибку #1 ", err)
			writer.WriteHeader(500)
			//401 — пользователь не авторизован.
			//500 — внутренняя ошибка сервера.
			return
		}
		fmt.Println("GetWithdrawals: получили данные", wdrls)
		//200 — успешная обработка запроса;
		//204 — нет ни одного списания.
		//200 — успешная обработка запроса.

		switch len(wdrls) {
		case 0: // отправляем ответ что нет ни одного списания
			fmt.Println("GetWithdrawals: не нашли ни одного списания", err)
			writer.WriteHeader(204)

		default:
			{
				writer.WriteHeader(200)
				writer.Header().Set("Content-Type", "application/json")

				strJSON, err := json.Marshal(wdrls)
				if err != nil {
					fmt.Println("GetWithdrawals: ушли в ошибку #2 ", err)
					writer.WriteHeader(500)
					return
				}

				if _, err = writer.Write(strJSON); err != nil {
					fmt.Println("GetWithdrawals: ушли в ошибку #3", err)
					writer.WriteHeader(500)
					return
				}
				fmt.Println("GetWithdrawals: все зашибись, отправили статус 200 и JSON ", strJSON)
			}
		}

	}
}
