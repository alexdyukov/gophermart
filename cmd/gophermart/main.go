package main

import (
	"log"
	"net/http"

	"github.com/alexdyukov/gophermart/internal/gophermart/domain/usecase"
	"github.com/alexdyukov/gophermart/internal/gophermart/handler"
	"github.com/alexdyukov/gophermart/internal/gophermart/repository/postgres"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {

	// Router
	gophermartRouter := chi.NewRouter()

	// Storage
	gophermartStore := postgres.NewGophermartStore()

	// Authentication handlers

	// Chi middlewares
	gophermartRouter.Use(middleware.Recoverer)
	// other middlewares, i.e. authorize

	// Handlers
	gophermartRouter.Post("/api/user/orders", handler.PostOrder(usecase.NewLoadOrderNumber(gophermartStore)))
	gophermartRouter.Get("/api/user/orders", handler.GetOrders(usecase.NewListOrderNums(gophermartStore)))
	gophermartRouter.Get("/api/user/balance", handler.GetBalance(usecase.NewShowBalanceState(gophermartStore)))
	gophermartRouter.Post("/api/user/balance/withdraw", handler.PostWithdraw(usecase.NewWithdrawFunds(gophermartStore)))
	gophermartRouter.Get("/api/user/balance/withdrawals", handler.GetWithdrawals(usecase.NewListWithdrawals(gophermartStore)))

	server := http.Server{
		Addr:    ":8089",
		Handler: gophermartRouter,
	}

	err := server.ListenAndServe()
	log.Print(err)
}
