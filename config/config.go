package config

import (
	"github.com/spf13/viper"
	"os"
)

type (
	Config struct {
		HttpServer HttpServer `mapstructure:"http_server" validate:"required"`
		Logger     Logger     `mapstructure:"logger" validate:"required"`
	}

	HttpServer struct {
		Host string `mapstructure:"host" validate:"required,ipv4"`
		Port int64  `mapstructure:"port" validate:"required"`
	}

	Logger struct {
		Level *int `mapstructure:"level" validate:"required"`
	}
)

func Load() (*Config, error) {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")

	path := os.Getenv("CONFIG")

	if len(path) != 0 {
		v.AddConfigPath(path)
	}

	v.AddConfigPath("config")

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	cfg := &Config{}

	err = v.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
