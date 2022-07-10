package main

import (
	"database/sql"
	"log"
	"net/http"

	authUsecase "github.com/alexdyukov/gophermart/internal/gophermart/auth/domain/usecase"
	"github.com/alexdyukov/gophermart/internal/gophermart/auth/gateway/token"
	authHandler "github.com/alexdyukov/gophermart/internal/gophermart/auth/handler"
	authPostgres "github.com/alexdyukov/gophermart/internal/gophermart/auth/repository/postgres"
	"github.com/alexdyukov/gophermart/internal/gophermart/config"
	"github.com/alexdyukov/gophermart/internal/gophermart/domain/usecase"
	"github.com/alexdyukov/gophermart/internal/gophermart/gateway/web"
	"github.com/alexdyukov/gophermart/internal/gophermart/handler"
	appMiddleware "github.com/alexdyukov/gophermart/internal/gophermart/handler/middleware"
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

	authStore, err := authPostgres.NewAuthStore(conn)
	if err != nil {
		log.Fatal(err)
	}

	appRouter := chi.NewRouter()
	appRouter.Use(chiMiddleware.Recoverer)

	var tokenTTL int64 = 3600 * 24 // to config
	jwtGateway := token.NewAuthJWTGateway(tokenTTL, []byte("secret"))
	appRouter.Post("/api/user/register", authHandler.RegisterPostHandler(
		authUsecase.NewRegisterUser(authStore, jwtGateway)))
	appRouter.Post("/api/user/login", authHandler.LoginPostHandler(authUsecase.NewLoginUser(authStore, jwtGateway)))

	appRouter.Group(func(subRouter chi.Router) {
		subRouter.Use(appMiddleware.Authentication(jwtGateway))
		subRouter.Post("/api/user/orders", handler.RegisterUserOrderPostHandler(
			usecase.NewLoadOrderNumber(gophermartStore, accrualGateway)))
		subRouter.Get("/api/user/orders", handler.ListUserOrdersGetHandler(
			usecase.NewListUserOrders(gophermartStore, accrualGateway)))
		subRouter.Get("/api/user/balance", handler.GetBalance(usecase.NewShowUserBalance(gophermartStore)))
		subRouter.Post("/api/user/balance/withdraw", handler.PostWithdraw(usecase.NewWithdrawUserFunds(gophermartStore)))
		subRouter.Get("/api/user/withdrawals", handler.GetWithdrawals(
			usecase.NewListUserWithdrawals(gophermartStore)))
	})

	server := http.Server{ // nolint:exhaustivestruct // ok, exhaustive // ok
		Addr:    addr,
		Handler: appRouter,
	}

	err = server.ListenAndServe()
	log.Print(err)
}
