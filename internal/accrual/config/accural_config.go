package config

import (
	"flag"
)

type ConfigAccrual struct {
	RunAddr   string
	DBConnect string
}

func NewAccrualConfig() *ConfigAccrual {
	var addr, databaseURI string

	flag.StringVar(&addr, "a", getEnv("ADDRESS", DefaultAddress), "Host IP address")
	flag.StringVar(&databaseURI, "d", getEnv("DATABASE_DSN", DefaultDB), "Connection string for DB")
	flag.Parse()

	cnf := ConfigAccrual{
		RunAddr:   addr,
		DBConnect: databaseURI,
	}

	return &cnf
}
