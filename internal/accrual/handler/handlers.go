package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/alexdyukov/gophermart/internal/accrual/domain/usecase"
	"github.com/go-chi/chi"
)

// GetOrders GET /api/orders/{number} — получение информации о расчёте начислений баллов лояльности;
func GetOrders(uc usecase.ShowLoyaltyPointsInputPort) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("GetOrders: запустился хендлер  /api/orders/{number}")
		number := chi.URLParam(request, "number")
		n, _ := strconv.Atoi(number)
		answ, err := uc.Execute(request.Context(), n)
		if err != nil {
			// 500 — внутренняя ошибка сервера.
			fmt.Println("GetOrders: ушли в ошибку #1 ", err)
			writer.WriteHeader(500)
			http.Error(writer, err.Error(), 500)
			return
		}

		// 429 — превышено количество запросов к сервису.

		// формируем ответ в нужном формате
		writer.WriteHeader(200)
		writer.Header().Set("Content-Type", "application/json")
		strJSON, err := json.Marshal(answ)

		_, err = writer.Write(strJSON)
		fmt.Println("GetOrders: все зашибись, отправили статус 200 и JSON ", string(strJSON))
	}
}

// PostOrders POST /api/orders — регистрация нового совершённого заказа
func PostOrders(uc usecase.CalculateLoyaltyPointsInputPort) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		dto := usecase.CalculateLoyaltyPointsInputDTO{}
		err := uc.Execute(request.Context(), dto)
		if err != nil {
			// todo: log error
			// todo: prepare response
		}
		//202 — заказ успешно принят в обработку;
		//400 — неверный формат запроса;
		//409 — заказ уже принят в обработку;
		//500 — внутренняя ошибка сервера.
		writer.WriteHeader(202)
	}
}

// PostGoods POST /api/goods — регистрация информации о новой механике вознаграждения за товар.
func PostGoods(uc usecase.RegisterMechanicInputPort) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// get manager from context
		actor := ""
		input := usecase.RegisterMechanicInputDTO{}
		err := uc.Execute(actor, input)
		if err != nil {
			// todo: log error
			// todo: prepare response
		}
		//200 — вознаграждение успешно зарегистрировано;
		//400 — неверный формат запроса;
		//409 — ключ поиска уже зарегистрирован;
		//500 — внутренняя ошибка сервера.
		writer.WriteHeader(200)
	}
}
