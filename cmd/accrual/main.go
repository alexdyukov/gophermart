package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/alexdyukov/gophermart/internal/accrual/config"
	"github.com/alexdyukov/gophermart/internal/accrual/domain/usecase"
	"github.com/alexdyukov/gophermart/internal/accrual/handler"
	"github.com/alexdyukov/gophermart/internal/accrual/repository/postgres"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	accrualConf := config.NewAccrualConfig()
	addr := accrualConf.RunAddr

	dsn := accrualConf.DBConnect

	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}

	accrualDB := postgres.NewAccrualDB(conn)

	// Router
	accrualRouter := chi.NewRouter()

	// Chi middlewares
	accrualRouter.Use(middleware.Recoverer)
	// other middlewares

	// Handlers
	accrualRouter.Get("/api/orders/{number}", handler.OrderCalculationGetHandler(
		usecase.NewShowOrderCalculation(accrualDB)))

	accrualRouter.Post("/api/orders", handler.RegisterOrderPostHandler(usecase.NewRegisterOrderReceipt(accrualDB)))

	accrualRouter.Post("/api/goods", handler.RegisterMechanicPostHandler(usecase.NewRegisterRewardMechanic(accrualDB)))

	server := http.Server{
		Addr:    addr,
		Handler: accrualRouter,
	}

	err = server.ListenAndServe()
	log.Print(err)
}
