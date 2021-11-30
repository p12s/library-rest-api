package config

import "github.com/kelseyhightower/envconfig"

// Config
type Config struct {
	Db     Db
	Server Server
	Logger Logger
}

// Db
type Db struct {
	Port     int    `envconfig:"DB_PORT" required:"true"`
	Host     string `envconfig:"DB_HOST" required:"true"`
	Driver   string `envconfig:"DB_DRIVER" required:"true"`
	User     string `envconfig:"DB_USER" required:"true"`
	Password string `envconfig:"DB_PASSWORD" required:"true"`
	SslMode  string `envconfig:"DB_SSL_MODE" required:"true"`
	Name     string `envconfig:"DB_NAME" required:"true"`
}

// Server
type Server struct {
	Port int `envconfig:"SERVER_PORT" required:"true"`
}

// Logger
type Logger struct {
	Host string `envconfig:"LOGGER_HOST" required:"true"`
}

// New - contructor
func New() (*Config, error) {
	cfg := new(Config)

	if err := envconfig.Process("db", &cfg.Db); err != nil {
		return nil, err
	}

	if err := envconfig.Process("server", &cfg.Server); err != nil {
		return nil, err
	}

	if err := envconfig.Process("logger", &cfg.Logger); err != nil {
		return nil, err
	}

	return cfg, nil
}
