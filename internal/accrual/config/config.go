package config

import (
	"os"
	"sync"
)

type AppConfig struct {
	RunAddr   string `env:"ACCRUAL_SYSTEM_ADDRESS"`
	DBConnect string `env:"DATABASE_URI"`
	sync.Once
}

// FlagGetter abstracts from flag source
type FlagGetter interface {
	Addr() string
	DatabaseURI() string
}

var appConfig *AppConfig = nil

func NewAppConfig() *AppConfig {
	if appConfig == nil {

		appFlags := parseFlags()

		appConfig = &AppConfig{}

		appConfig.Do(func() {
			appConfig.configure(&appFlags)
		})
		return appConfig
	}
	return appConfig
}

func (a *AppConfig) configure(appFlags FlagGetter) {

	a.RunAddr = getEnv("ACCRUAL_SYSTEM_ADDRESS", DefaultAddressAS)
	a.DBConnect = getEnv("DATABASE_URI", DefaultDB)

	if appFlags.Addr() != "" {
		a.RunAddr = appFlags.Addr()
	}
	if appFlags.DatabaseURI() != "" {
		a.DBConnect = appFlags.DatabaseURI()
	}
}

func getEnv(key string, defaultVal string) string {

	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal

}
