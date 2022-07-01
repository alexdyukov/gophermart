package main

import (
	"log"
	"net/http"

	"github.com/alexdyukov/gophermart/internal/accrual/domain/usecase"
	"github.com/alexdyukov/gophermart/internal/accrual/handler"
	"github.com/alexdyukov/gophermart/internal/accrual/repository/memory"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	// Router
	accrualRouter := chi.NewRouter()

	// Chi middlewares
	accrualRouter.Use(middleware.Recoverer)
	// other middlewares

	// Storage
	memRepo := memory.NewAccrualStore()

	// Handlers
	accrualRouter.Get("/api/orders/{number}", handler.GetOrders(usecase.NewShowLoyaltyPoints(memRepo)))
	accrualRouter.Post("/api/orders", handler.PostOrders(usecase.NewCalculateLoyaltyPoints(memRepo)))
	accrualRouter.Post("/api/goods", handler.PostGoods(usecase.NewRegisterMechanic(memRepo)))

	server := http.Server{ // nolint:exhaustivestruct // ok
		Addr:    ":8088",
		Handler: accrualRouter,
	}

	err := server.ListenAndServe()
	log.Print(err)
}
