package config

import (
	"flag"
	"os"
)

type ConfigGophermart struct {
	RunAddr       string
	DBConnect     string
	AccSystemAddr string
}

const (
	DefaultAddress   = "127.0.0.1:8088"
	DefaultDB        = "" //host=localhost dbname=ya_pr_devops
	DefaultAddressAS = "127.0.0.1:8089"
)

func NewGophermartConfig() *ConfigGophermart {
	var addr, databaseURI string
	var addrAccrualSystem string

	flag.StringVar(&addr, "a", getEnv("RUN_ADDRESS", DefaultAddress), "Host IP address")
	flag.StringVar(&databaseURI, "d", getEnv("DATABASE_URI", DefaultDB), "Connection string for DB")
	flag.StringVar(&addrAccrualSystem, "r", getEnv("ACCRUAL_SYSTEM_ADDRESS", DefaultAddressAS), "Host IP address accrual system")
	flag.Parse()

	return &ConfigGophermart{
		RunAddr:       addr,
		DBConnect:     databaseURI,
		AccSystemAddr: addrAccrualSystem,
	}
}

func getEnv(key string, defaultVal string) string {

	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
