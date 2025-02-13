package config

import (
	"os"
	"reflect"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	Auth       AuthConfig
	Redis      RedisConfig
	PostgreSQL PostgreSQLConfig
	Server     ServerConfig
	Logger     LoggerConfig
	Storage    StorageConfig
}
type PostgreSQLConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSL      bool
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
	TLS          bool
}

type ServerConfig struct {
	AppVersion        string
	Port              string
	Mode              string
	ReadTimeout       int
	WriteTimeout      int
	SSL               bool
	CtxDefaultTimeout int
	SuperAdmin        string
}

type AuthConfig struct {
	Secret   string
	Expire   int
	Issuer   string
	Audience string
}

type LoggerConfig struct {
	Level string
	Mode  string
}

type StorageConfig struct {
	Host        string
	Container   string
	AccountName string
	AccountKey  string
}

func GetConfig() *Config {
	var c Config
	vp := viper.New()
	appEnv := os.Getenv("APP_ENV")

	switch appEnv {
	case "heroku":
		envs := os.Environ()
		mapEnv := map[string]interface{}{}
		for _, env := range envs {
			key, value := strings.Split(env, "=")[0], strings.Split(env, "=")[1]
			mapEnv[key] = value
		}
		mapEnv["SERVER.PORT"] = os.Getenv("PORT")
		vp.MergeConfigMap(mapEnv)
		bindEnvs(vp, c)
	default:
		vp.SetConfigName("config")
		vp.SetConfigType("yaml")
		vp.AddConfigPath("./config")
		if err := vp.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				log.Fatal().Msg("Config file not found")
			}
		}
	}
	if err := vp.Unmarshal(&c); err != nil {
		log.Fatal().Msg("Unable to unmarshal config")
	}
	return &c
}

func bindEnvs(v *viper.Viper, iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		fieldv := ifv.Field(i)
		t := ift.Field(i)
		name := strings.ToLower(t.Name)
		tag, ok := t.Tag.Lookup("mapstructure")
		if ok {
			name = tag
		}
		path := append(parts, name)
		switch fieldv.Kind() {
		case reflect.Struct:
			bindEnvs(v, fieldv.Interface(), path...)
		default:
			v.BindEnv(strings.Join(path, "."))
		}
	}
}
