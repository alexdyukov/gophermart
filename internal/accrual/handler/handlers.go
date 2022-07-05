package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/alexdyukov/gophermart/internal/accrual/domain/usecase"
	"github.com/go-chi/chi"
)

// OrderCalculationGetHandler GET /api/orders/{number} — получение информации о расчёте начислений баллов лояльности;
// 429 — превышено количество запросов к сервису.
// 500 — внутренняя ошибка сервера.
func OrderCalculationGetHandler(showLoyaltyUsecase usecase.ShowOrderCalculationPrimaryPort) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		fmt.Println("ping get orders")

		number := chi.URLParam(request, "number")

		n, err := strconv.Atoi(number)
		if err != nil {
			log.Println(err)
		}

		err = showLoyaltyUsecase.Execute(n)
		if err != nil {
			log.Println(err)
		}

		writer.WriteHeader(http.StatusOK)
	}
}

// RegisterOrderPostHandler POST /api/orders — регистрация нового совершённого заказа.
// 202 — заказ успешно принят в обработку;
// 400 — неверный формат запроса;
// 409 — заказ уже принят в обработку;
// 500 — внутренняя ошибка сервера.
func RegisterOrderPostHandler(calculateLoyaltyUsecase usecase.RegisterPurchasedOrderPrimaryPort) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		dto := usecase.RegisterPurchasedOrderInputDTO{}

		err := calculateLoyaltyUsecase.Execute(dto)
		if err != nil {
			log.Println(err)
		}

		writer.WriteHeader(http.StatusOK)
	}
}

// RegisterMechanicPostHandler POST /api/goods — регистрация информации о новой механике вознаграждения за товар.
// 200 — вознаграждение успешно зарегистрировано;
// 400 — неверный формат запроса;
// 409 — ключ поиска уже зарегистрирован;
// 500 — внутренняя ошибка сервера.
func RegisterMechanicPostHandler(registerMechanicUsecase usecase.RegisterMechanicPrimaryPort) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		actor := ""
		input := usecase.RegisterMechanicInputDTO{}

		err := registerMechanicUsecase.Execute(actor, &input)
		if err != nil {
			log.Println(err)

			return
		}

		writer.WriteHeader(http.StatusOK)
	}
}
