package config

import (
	"flag"
	"os"
)

const (
	DefaultAddress = "127.0.0.1:8088"
	DefaultDB      = "" //host=localhost dbname=ya_pr_devops
)

type ConfigAccrual struct {
	RunAddr   string
	DBConnect string
}

func NewAccrualConfig() *ConfigAccrual {
	var addr, databaseURI string

	flag.StringVar(&addr, "a", getEnv("RUN_ADDRESS", DefaultAddress), "Host IP address")
	flag.StringVar(&databaseURI, "d", getEnv("DATABASE_URI", DefaultDB), "Connection string for DB")
	flag.Parse()

	return &ConfigAccrual{
		RunAddr:   addr,
		DBConnect: databaseURI,
	}
}

func getEnv(key string, defaultVal string) string {

	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal

}
