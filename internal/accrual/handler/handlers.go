package handler

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/alexdyukov/gophermart/internal/accrual/domain/usecase"
	"github.com/go-chi/chi"
)

// OrderCalculationGetHandler GET /api/orders/{number} — получение информации о расчёте начислений баллов лояльности;
// 429 — превышено количество запросов к сервису.
// 500 — внутренняя ошибка сервера.
func OrderCalculationGetHandler(showOrderCalculationUsecase usecase.ShowOrderCalculationPrimaryPort) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		number := chi.URLParam(request, "number")

		output, err := showOrderCalculationUsecase.Execute(request.Context(), number)
		if err != nil {
			if errors.Is(err, usecase.ErrIncorrectOrderNumber) {
				writer.WriteHeader(http.StatusBadRequest)

				return
			}

			writer.WriteHeader(http.StatusInternalServerError)

			return
		}

		result, err := json.Marshal(output)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)

			return
		}

		writer.Header().Set("Content-Type", "application/json")
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
func RegisterOrderPostHandler(
	registerPurchasedOrderUsecase usecase.RegisterOrderReceiptPrimaryPort,
	calculateRewardUsecase usecase.CalculateRewardPrimaryPort,
) http.HandlerFunc { // nolint:whitespace // ok
	return func(writer http.ResponseWriter, request *http.Request) {
		bytes, err := io.ReadAll(request.Body)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)

			return
		}

		orderReceiptDTO := usecase.RegisterOrderReceiptInputDTO{}

		err = json.Unmarshal(bytes, &orderReceiptDTO)
		if err != nil {
			log.Println("cant unmarshal", err)
			writer.WriteHeader(http.StatusBadRequest)

			return
		}

		orderReceipt, err := registerPurchasedOrderUsecase.Execute(request.Context(), &orderReceiptDTO)
		if err != nil {
			if errors.Is(err, usecase.ErrIncorrectOrderNumber) {
				log.Println("incorrect order number", err)
				writer.WriteHeader(http.StatusBadRequest)

				return
			}

			if errors.Is(err, usecase.ErrOrderAlreadyExist) {
				writer.WriteHeader(http.StatusConflict)

				return
			}

			writer.WriteHeader(http.StatusInternalServerError)

			return
		}

		err = calculateRewardUsecase.Execute(request.Context(), orderReceipt)
		if err != nil {
			log.Println(err)
		}

		writer.WriteHeader(http.StatusAccepted)
	}
}

// RegisterMechanicPostHandler POST /api/goods — регистрация информации о новой механике вознаграждения за товар.
// 200 — вознаграждение успешно зарегистрировано;
// 400 — неверный формат запроса;
// 409 — ключ поиска уже зарегистрирован;
// 500 — внутренняя ошибка сервера.
func RegisterMechanicPostHandler(registerRewardUsecase usecase.RegisterRewardMechanicPrimaryPort) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		bytes, err := io.ReadAll(request.Body)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)

			return
		}

		registerRewardInputDTO := usecase.RegisterRewardMechanicInputDTO{}

		err = json.Unmarshal(bytes, &registerRewardInputDTO)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)

			return
		}

		err = registerRewardUsecase.Execute(request.Context(), &registerRewardInputDTO)
		if err != nil {
			if errors.Is(err, usecase.ErrRewardAlreadyExists) {
				writer.WriteHeader(http.StatusConflict)

				return
			}

			writer.WriteHeader(http.StatusBadRequest)

			return
		}

		writer.WriteHeader(http.StatusOK)
	}
}
