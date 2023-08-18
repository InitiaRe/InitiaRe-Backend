package config

import (
	"os"
	"reflect"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	Auth       AuthConfig       `mapstructure:"auth"`
	Redis      RedisConfig      `mapstructure:"redis"`
	PostgreSQL PostgreSQLConfig `mapstructure:"postgresql"`
	Server     ServerConfig     `mapstructure:"server"`
	Logger     LoggerConfig     `mapstructure:"logger"`
	Storage    StorageConfig    `mapstructure:"storage"`
}
type PostgreSQLConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	PoolSize     int    `mapstructure:"poolsize"`
	PoolTimeout  int    `mapstructure:"pooltimeout"`
	MinIdleConns int    `mapstructure:"minidleconns"`
	DB           int    `mapstructure:"db"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
}

type ServerConfig struct {
	AppVersion        string `mapstructure:"app_version"`
	Port              string `mapstructure:"port"`
	Mode              string `mapstructure:"mode"`
	ReadTimeout       int    `mapstructure:"readtimeout"`
	WriteTimeout      int    `mapstructure:"writetimeout"`
	SSL               bool   `mapstructure:"ssl"`
	CtxDefaultTimeout int    `mapstructure:"ctxdefaulttimeout"`
}

type AuthConfig struct {
	Secret   string `mapstructure:"secret"`
	Expire   int    `mapstructure:"expire"`
	Issuer   string `mapstructure:"issuer"`
	Audience string `mapstructure:"audience"`
}

type LoggerConfig struct {
	Level string `mapstructure:"level"`
	Mode  string `mapstructure:"mode"`
}

type StorageConfig struct {
	Host        string `mapstructure:"host"`
	Container   string `mapstructure:"container"`
	AccountName string `mapstructure:"accountname"`
	AccountKey  string `mapstructure:"accountkey"`
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
