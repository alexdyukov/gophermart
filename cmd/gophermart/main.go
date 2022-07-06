package main

import (
	"context"
	"fmt"
	"github.com/alexdyukov/gophermart/internal/gophermart/config"
	"github.com/alexdyukov/gophermart/internal/gophermart/domain/usecase"
	"github.com/alexdyukov/gophermart/internal/gophermart/handler"
	"github.com/alexdyukov/gophermart/internal/gophermart/repository/postgres"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
)

func main() {
	//config
	appConf := config.NewGophermartConfig()
	// Router
	gophermartRouter := chi.NewRouter()

	// Storage

	//gophermartStore := memory.NewGophermartStore()
	//gophermartStore := ""

	//if appConf.DBConnect != "" {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dataBase, err := postgres.OpenDB(appConf.DBConnect)

	if err != nil {
		fmt.Println("ошибка при открытии БД", err)
		return
	}

	defer func() { _ = dataBase.Close() }()

	if err := postgres.InitSchema(ctx, dataBase); err != nil {
		fmt.Println("ошибка при создании инициализации схемы", err)
	}

	gophermartStore := postgres.NewGophermartStore(dataBase)

	// пробуем записать пользователя в таблицу
	idUser := "057f2f06-9e6d-4cf2-aa77-7f4cc1a51f9b"
	gophermartStore.SaveUser(ctx, "Oesya", "Olesya", idUser)

	gophermartStore.SaveOrderTest(ctx, idUser, "9278923470", 500)
	gophermartStore.SaveOrderTest(ctx, idUser, "12345678903", 324.82)

	//}

	// Authentication handlers

	// Chi middlewares
	gophermartRouter.Use(middleware.Recoverer)
	// other middlewares, i.e. authorize

	// Handlers
	gophermartRouter.Post("/api/user/orders", handler.PostOrder(usecase.NewLoadOrderNumber(gophermartStore)))
	gophermartRouter.Get("/api/user/orders", handler.GetOrders(usecase.NewListOrderNums(gophermartStore)))
	gophermartRouter.Get("/api/user/balance", handler.GetBalance(usecase.NewShowBalanceState(gophermartStore)))
	gophermartRouter.Post("/api/user/balance/withdraw", handler.PostWithdraw(usecase.NewWithdrawFunds(gophermartStore)))
	gophermartRouter.Get("/api/user/withdrawals", handler.GetWithdrawals(usecase.NewListWithdrawals(gophermartStore))) // BeOl - видать это ошибка

	server := http.Server{
		Addr:    appConf.RunAddr,
		Handler: gophermartRouter,
	}

	err = server.ListenAndServe()
	log.Print(err)
}
