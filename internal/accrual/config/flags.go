package config

import (
	"flag"
)

type AppFlags struct {
	addr        *string // -a
	databaseURI *string // -d
}

func (p *AppFlags) Addr() string {
	return *p.addr
}

func (p *AppFlags) DatabaseURI() string {
	return *p.databaseURI
}

const (
	DefaultAddressAS = "127.0.0.1:8089"
	DefaultDB        = "" //host=localhost dbname=ya_pr_devops
)

func parseFlags() AppFlags {
	var addrAccrualSystem, databaseURI string

	flag.StringVar(&addrAccrualSystem, "r", DefaultAddressAS, "Host IP address accrual system")
	flag.StringVar(&databaseURI, "d", DefaultDB, "Connection string for DB")
	flag.Parse()

	return AppFlags{addr: &addrAccrualSystem, databaseURI: &databaseURI}

}
