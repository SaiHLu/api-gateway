package config

import "github.com/caarlos0/env/v11"

type EnvConfig struct {
	Port       int              `env:"port"`
	Mode       string           `env:"mode"`
	GrpcServer GrpcServerConfig `env:"grpc_server"`
	Cors       CorsConfig       `env:"cors"`
}

type GrpcServerConfig struct {
	ProductService string `env:"product_service"`
	OrderService   string `env:"order_service"`
}

type CorsConfig struct {
	AllowOrigins     []string `env:"allow_origins"`
	AllowMethods     []string `env:"allow_methods"`
	AllowHeaders     []string `env:"allow_headers"`
	AllowCredentials bool     `env:"allow_credentials"`
	ExposeHeaders    []string `env:"expose_headers"`
}

func LoadConfig() (*EnvConfig, error) {
	var cfg EnvConfig
	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
