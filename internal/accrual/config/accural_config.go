package config

import (
	"flag"
	"os"
)

type Accrual struct {
	RunAddr   string
	DBConnect string
}

const (
	DefaultAddress = "127.0.0.1:8088"
	DefaultDB      = ""
)

func NewAccrualConfig() *Accrual {
	var addr, databaseURI string

	flag.StringVar(&addr, "a", getEnv("RUN_ADDRESS", DefaultAddress), "Host IP address")
	flag.StringVar(&databaseURI, "d", getEnv("DATABASE_URI", DefaultDB), "Connection string for DB")
	flag.Parse()

	return &Accrual{
		RunAddr:   addr,
		DBConnect: databaseURI,
	}
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
