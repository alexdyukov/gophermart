package config

import (
	"flag"
	"os"
)

type Gophermart struct {
	RunAddr       string
	DBConnect     string
	AccSystemAddr string
}

const (
	DefaultAddress   = "127.0.0.1:8089"
	DefaultDB        = ""
	DefaultAddressAS = "127.0.0.1:8088"
)

func NewGophermartConfig() *Gophermart {
	var addr string

	var databaseURI string

	var addrAccrualSystem string

	flag.StringVar(&addr, "a", getEnv("RUN_ADDRESS", DefaultAddress), "Host IP address")
	flag.StringVar(&databaseURI, "d", getEnv("DATABASE_URI", DefaultDB), "Connection string for DB")
	flag.StringVar(&addrAccrualSystem, "r",
		getEnv("ACCRUAL_SYSTEM_ADDRESS", DefaultAddressAS), "Host IP address accrual system")
	flag.Parse()

	return &Gophermart{
		RunAddr:       addr,
		DBConnect:     databaseURI,
		AccSystemAddr: addrAccrualSystem,
	}
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
