package main

import (
	"fmt"
	"github.com/alexdyukov/gophermart/internal/accrual/config"
	"log"
	"net/http"

	"github.com/alexdyukov/gophermart/internal/accrual/domain/usecase"
	"github.com/alexdyukov/gophermart/internal/accrual/handler"
	"github.com/alexdyukov/gophermart/internal/accrual/repository/memory"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {

	// config..
	appConf := config.NewAppConfig()
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

	fmt.Printf("#Run accrual with IP: %s \n", appConf.RunAddr)
	server := http.Server{
		Addr:    appConf.RunAddr,
		Handler: accrualRouter,
	}
	err := server.ListenAndServe()
	log.Print(err)
}
