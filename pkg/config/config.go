package config

import (
	"github.com/jessevdk/go-flags"
)

var config *Config

func GetConfig() *Config {
	if nil == config {
		config = &Config{}
		parser := flags.NewParser(config, flags.Default)
		if _, err := parser.Parse(); err != nil {
			panic(err)
		}
	}

	return config
}

type Config struct {
	AppHost string `short:"h" long:"app-host" env:"APP_HOST" description:"Application server host"`
	AppPort int    `short:"p" long:"app-port" env:"APP_PORT" description:"Application server port"`

	MySQLDBO string `long:"mysql-dbo" env:"MYSQL_DBO" required:"true"`
	RedisAddr string `long:"redis-addr" env:"REDIS_ADDR" required:"true"`
	Secret string `long:"secret" env:"SECRET" required:"true"`
	ApiKey string `long:"api_key" env:"API_KEY" required:"true"`
}
