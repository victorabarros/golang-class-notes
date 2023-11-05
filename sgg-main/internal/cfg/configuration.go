package cfg

import "time"

type Configuration struct {
	Prefix string `env:"PREFIX" env-default:"/api/v1"`
	Cron   struct {
		DeleteArticles string `env:"CRON_DELETE_JOBS" env-default:"0 0 0 * * *"` // kicks-off at midnight everyday
	}
	HTTP struct {
		Port           int           `env:"HTTP_PORT" env-default:"8080"`
		DefaultTimeout time.Duration `env:"DEFAULT_TIMEOUT" env-default:"30s"`
	}
	GRPC struct {
		Port int `env:"GRPC_PORT" env-default:"9090"`
	}
}
