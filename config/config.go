package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Auth       AuthConfig
	Redis      RedisConfig
	PostgreSQL PostgreSQLConfig
	Server     ServerConfig
}
type PostgreSQLConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

type RedisConfig struct {
	Host         string
	Port         int
	PoolSize     int
	PoolTimeout  int
	MinIdleConns int
	DB           int
	Username     string
	Password     string
}

type ServerConfig struct {
	AppVersion        string
	Port              string
	Mode              string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	SSL               bool
	CtxDefaultTimeout time.Duration
	Debug             bool
}

type AuthConfig struct {
	JWTSecret string
	Expire    int
	Issuer    string
	Audience  string
}

func GetConfig() *Config {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("Config file not found")
		}
	}
	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		panic("Unable to unmarshal config")
	}
	return &c
}
