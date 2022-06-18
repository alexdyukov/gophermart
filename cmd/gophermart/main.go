package main

import (
	"log"
	"net/http"

	"github.com/alexdyukov/gophermart/internal/config"
	svc "github.com/alexdyukov/gophermart/internal/gophermartsvc"
	"github.com/alexdyukov/gophermart/internal/handler"
	"github.com/alexdyukov/gophermart/internal/storage"
)

func main() {
	conf := config.Get()

	stor := storage.New(conf.StorageType)
	svc := svc.New(conf.AccrualAddress, stor)
	h := handler.New(conf, svc)

	log.Fatal(http.ListenAndServe(conf.ServerAddress.String(), h))
}
