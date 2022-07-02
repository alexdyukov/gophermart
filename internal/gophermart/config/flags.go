package config

import (
	"flag"
)

type AppFlags struct {
	addr              *string // -a
	databaseURI       *string // -d
	addrAccrualSystem *string // -r
}

func (p *AppFlags) Addr() string {
	return *p.addr
}

func (p *AppFlags) DatabaseURI() string {
	return *p.databaseURI
}

func (p *AppFlags) AddrAccuralSystem() string {
	return *p.addrAccrualSystem
}

const (
	DefaultAddress   = "127.0.0.1:41343"
	DefaultDB        = "" //host=localhost dbname=ya_pr_devops
	DefaultAddressAS = "127.0.0.1:44157"
)

func parseFlags() AppFlags {
	var addr, databaseURI string
	var addrAccrualSystem string

	flag.StringVar(&addr, "a", DefaultAddress, "Host IP address")
	flag.StringVar(&databaseURI, "d", DefaultDB, "Connection string for DB")
	flag.StringVar(&addrAccrualSystem, "r", DefaultAddressAS, "Host IP address accrual system")
	flag.Parse()

	return AppFlags{addr: &addr, databaseURI: &databaseURI, addrAccrualSystem: &addrAccrualSystem}

}
