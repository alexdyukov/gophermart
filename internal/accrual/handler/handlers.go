package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alexdyukov/gophermart/internal/accrual/domain/usecase"
	"github.com/go-chi/chi"
)

// GetOrders GET /api/orders/{number} — получение информации о расчёте начислений баллов лояльности;
func GetOrders(uc usecase.ShowLoyaltyPointsInputPort) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("bla")
		number := chi.URLParam(request, "number")
		n, _ := strconv.Atoi(number)
		err := uc.Execute(n)
		if err != nil {
			// todo: log error
			// todo: prepare response
		}
		// 429 — превышено количество запросов к сервису.
		// 500 — внутренняя ошибка сервера.
		writer.WriteHeader(200)
	}
}

// PostOrders POST /api/orders — регистрация нового совершённого заказа
func PostOrders(uc usecase.CalculateLoyaltyPointsInputPort) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		dto := usecase.CalculateLoyaltyPointsInputDTO{}
		err := uc.Execute(dto)
		if err != nil {
			// todo: log error
			// todo: prepare response
		}
		//202 — заказ успешно принят в обработку;
		//400 — неверный формат запроса;
		//409 — заказ уже принят в обработку;
		//500 — внутренняя ошибка сервера.
		writer.WriteHeader(200)
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
