package config

import "flag"

type ConfigGophermart struct {
	RunAddr       string
	DBConnect     string
	AccSystemAddr string
}

func NewGophermartConfig() *ConfigGophermart {
	var addr, databaseURI string
	var addrAccrualSystem string

	flag.StringVar(&addr, "a", getEnv("RUN_ADDRESS", DefaultAddress), "Host IP address")
	flag.StringVar(&databaseURI, "d", getEnv("DATABASE_URI", DefaultDB), "Connection string for DB")
	flag.StringVar(&addrAccrualSystem, "r", getEnv("ACCRUAL_SYSTEM_ADDRESS", DefaultAddressAS), "Host IP address accrual system")

	flag.Parse()

	cnf := ConfigGophermart{
		RunAddr:       addr,
		DBConnect:     databaseURI,
		AccSystemAddr: addrAccrualSystem,
	}

	return &cnf
}
