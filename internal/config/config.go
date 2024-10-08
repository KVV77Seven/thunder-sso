package config

import "time"

type Config struct {
	Env             string        `yaml:"env" env-required:"true"`
	StoragePath     string        `yaml:"storage_path" env-required:"true"`
	AccessTokenTTL  time.Duration `yaml:"access_token_ttl" env-required:"true"`
	RefreshTokenTTL time.Duration `yaml:"refresh_token_ttl" env-required:"true"`
	GRPC            GRPCConfig    `yaml:"grpc" env-required:"true"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port" env-required:"true"`
	Timeout time.Duration `yaml:"timeout" env-required:"true"`
}
