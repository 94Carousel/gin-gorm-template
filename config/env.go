package config

import "github.com/caarlos0/env"

// Config defined global config
type Config struct {
	Addr          string `env:"GIN_ADDR" envDefault:"localhost:3000"`
	Port          int    `env:"PORT" envDefault:"3000"`
	RunMode       string `env:"GIN_MODE" envDefault:"debug"`
	AppName       string `env:"GIN_APP_NAME" envDefault:"gin"`
	SecretKey     string `env:"GIN_SECRET_KEY"`
	DBAdapter     string `env:"DB_ADAPTER" envDefault:"sqlite"`
	DBHostName    string `env:"DB_HOSTNAME" envDefault:"localhost"`
	DBDataBase    string `env:"DB_DATABASE" envDefault:"gin.db"`
	DBUserName    string `env:"DB_USERNAME"`
	DBPassWord    string `env:"DB_PASSWORD"`
	DBPort        string `env:"DB_PORT"`
	RedisAddr     string `env:"REDIS_ADDR" envDefault:"localhost:6379"`
	RedisPassWord string `env:"REDIS_PASSWORD" envDefault:""`
	RedisDB       int    `env:"REDIS_DB" envDefault:"0"`
}

// EnvConfig Config
var EnvConfig *Config

// InitEnv return global config
func InitEnv() *Config {
	cfg := Config{}
	env.Parse(&cfg)
	EnvConfig = &cfg
	return &cfg
}
