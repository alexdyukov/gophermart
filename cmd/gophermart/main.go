package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
	"log"
	"net/http"
	"time"

	"github.com/alexdyukov/gophermart/internal/gophermart/config"
	"github.com/alexdyukov/gophermart/internal/gophermart/domain/usecase"
	"github.com/alexdyukov/gophermart/internal/gophermart/gateway/web"
	"github.com/alexdyukov/gophermart/internal/gophermart/handler"
	"github.com/alexdyukov/gophermart/internal/gophermart/repository/postgres"
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() { // nolint:funlen // ok
	// Configure application
	gophermartConf := config.NewGophermartConfig()
	dbConnString := gophermartConf.DBConnect
	addr := gophermartConf.RunAddr

	// sub service
	accrualAddr := gophermartConf.AccSystemAddr
	gatewayEndpoint := "/api/orders/"
	accrualGateway := web.NewAccrualGateway(accrualAddr, gatewayEndpoint) // to config

	conn, err := sql.Open("pgx", dbConnString)
	if err != nil {
		log.Fatal(err)
	}

	gophermartStore, err := postgres.NewGophermartDB(conn)
	if err != nil {
		log.Fatal(err)
	}

	//authStore, err := authPostgres.NewAuthStore(conn)
	if err != nil {
		log.Fatal(err)
	}

	idUser := "057f2f06-9e6d-4cf2-aa77-7f4cc1a51f9b"
	//	gophermartStore.SaveUser(ctx, "Oesya", "Olesya", idUser)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err = gophermartStore.SaveOrderTest(ctx, idUser, 888355533, 350, sharedkernel.PROCESSED, time.Date(2019, time.May, 15, 17, 45, 12, 0, time.Local))
	err = gophermartStore.SaveOrderTest(ctx, idUser, 324235553, 200, sharedkernel.NEW, time.Date(2022, time.May, 15, 17, 45, 12, 0, time.Local))
	err = gophermartStore.SaveOrderTest(ctx, idUser, 988355568, 200, sharedkernel.PROCESSING, time.Date(2022, time.May, 15, 17, 45, 12, 0, time.Local))
	err = gophermartStore.SaveOrderTest(ctx, idUser, 104535323, 200, sharedkernel.PROCESSED, time.Date(2018, time.May, 15, 17, 45, 12, 0, time.Local))

	if err != nil {
		fmt.Println("ошибка при записи заказа ", err)
		return
	}
	appRouter := chi.NewRouter()
	appRouter.Use(chiMiddleware.Recoverer)

	//var tokenTTL int64 = 3600 * 24 // to config
	//jwtGateway := token.NewAuthJWTGateway(tokenTTL, []byte("secret"))
	//appRouter.Post("/api/user/register", authHandler.RegisterPostHandler(
	//	authUsecase.NewRegisterUser(authStore, jwtGateway)))
	//appRouter.Post("/api/user/login", authHandler.LoginPostHandler(authUsecase.NewLoginUser(authStore, jwtGateway)))

	appRouter.Group(func(subRouter chi.Router) {
		//subRouter.Use(appMiddleware.Authentication(jwtGateway))
		subRouter.Post("/api/user/orders", handler.RegisterUserOrderPostHandler(
			usecase.NewLoadOrderNumber(gophermartStore, accrualGateway)))
		subRouter.Get("/api/user/orders", handler.ListUserOrdersGetHandler(usecase.NewListUserOrders(gophermartStore)))
		subRouter.Get("/api/user/balance", handler.GetBalance(usecase.NewShowUserBalance(gophermartStore)))
		subRouter.Post("/api/user/balance/withdraw", handler.PostWithdraw(usecase.NewWithdrawUserFunds(gophermartStore)))
		subRouter.Get("/api/user/balance/withdrawals", handler.GetWithdrawals(
			usecase.NewListUserWithdrawals(gophermartStore)))
	})

	server := http.Server{ // nolint:exhaustivestruct // ok, exhaustive // ok
		Addr:    addr,
		Handler: appRouter,
	}

	err = server.ListenAndServe()
	log.Print(err)
}
