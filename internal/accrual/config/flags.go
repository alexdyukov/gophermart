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
	DefaultAddress = "127.0.0.1:8088"
	DefaultDB      = "" //host=localhost dbname=ya_pr_devops
)

func parseFlags() AppFlags {
	var addr, databaseURI string

	flag.StringVar(&addr, "a", DefaultAddress, "Host IP address")
	flag.StringVar(&databaseURI, "d", DefaultDB, "Connection string for DB")
	flag.Parse()

	return AppFlags{addr: &addr, databaseURI: &databaseURI}

}
