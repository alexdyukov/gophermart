package main

import (
	"log"
	"net/http"

	svc "github.com/alexdyukov/gophermart/internal/accrualsvc"
	"github.com/alexdyukov/gophermart/internal/config"
	"github.com/alexdyukov/gophermart/internal/handler"
	"github.com/alexdyukov/gophermart/internal/storage"
)

func main() {
	conf := config.Get()

	stor := storage.New(conf.StorageType)
	svc := svc.New(stor)
	h := handler.New(conf, svc)

	log.Fatal(http.ListenAndServe(conf.ServerAddress.String(), h))
}
