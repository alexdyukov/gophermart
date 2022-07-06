package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/alexdyukov/gophermart/internal/accrual/domain/usecase"
	"github.com/go-chi/chi"
)

// OrderCalculationGetHandler GET /api/orders/{number} — получение информации о расчёте начислений баллов лояльности;
// 429 — превышено количество запросов к сервису.
// 500 — внутренняя ошибка сервера.
func OrderCalculationGetHandler(showOrderCalculationUsecase usecase.ShowOrderCalculationPrimaryPort) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		n := chi.URLParam(request, "number")

		number, err := strconv.Atoi(n)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)

			return
		}

		output, err := showOrderCalculationUsecase.Execute(request.Context(), number)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError) //500

			return
		}

		result, err := json.Marshal(output)
		if err != nil {
			log.Println(err)
		}

		writer.WriteHeader(http.StatusOK)

		_, err = writer.Write(result)
		if err != nil {
			log.Println(err)
		}
	}
}

// RegisterOrderPostHandler POST /api/orders — регистрация нового совершённого заказа.
// 202 — заказ успешно принят в обработку;
// 400 — неверный формат запроса;
// 409 — заказ уже принят в обработку;
// 500 — внутренняя ошибка сервера.
func RegisterOrderPostHandler(registerPurchasedOrderUsecase usecase.RegisterOrderReceiptPrimaryPort) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		bytes, err := io.ReadAll(request.Body)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)

			return
		}

		orderReceiptDTO := usecase.RegisterOrderReceiptInputDTO{}

		err = json.Unmarshal(bytes, &orderReceiptDTO)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)

			return
		}

		err = registerPurchasedOrderUsecase.Execute(orderReceiptDTO)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
		}

		writer.WriteHeader(http.StatusOK)
	}
}

// RegisterMechanicPostHandler POST /api/goods — регистрация информации о новой механике вознаграждения за товар.
// 200 — вознаграждение успешно зарегистрировано;
// 400 — неверный формат запроса;
// 409 — ключ поиска уже зарегистрирован;
// 500 — внутренняя ошибка сервера.
func RegisterMechanicPostHandler(registerRewardUsecase usecase.RegisterRewardMechanicPrimaryPort) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		actor := ""
		input := usecase.RegisterRewardMechanicInputDTO{}

		err := registerRewardUsecase.Execute(actor, &input)
		if err != nil {
			log.Println(err)

			return
		}

		writer.WriteHeader(http.StatusOK)
	}
}
