package config

import (
	"context"
	"os"
	"reflect"
	"strings"

	"github.com/go-playground/validator"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type (
	DB struct {
		ConnUri string `mapstructure:"DB_CONN_URI" validate:"required"`
	}

	JWT struct {
		SigningKey string `mapstructure:"JWT_SIGNING_KEY" validate:"required"`
	}

	Configuration struct {
		DB  DB  `mapstructure:",squash"`
		JWT JWT `mapstructure:",squash"`
	}
)

func InitConfig(ctx context.Context) *Configuration {
	var cfg Configuration

	viper.SetConfigType("env")
	envFile := os.Getenv("ENV_FILE")
	if envFile == "" {
		envFile = ".env"
	}

	viper.AutomaticEnv()
	BindEnvs(&cfg, "")

	_, err := os.Stat(envFile)
	if !os.IsNotExist(err) {
		viper.SetConfigFile(envFile)

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("failed to read config: %v", err)
		}
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("failed to bind config: %v", err)
	}

	validate := validator.New()
	if err := validate.Struct(cfg); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			log.Errorf("invalid config: %v\n", err)
		}
		log.Fatal("failed to load config")
	}

	return &cfg
}

func BindEnvs(cfg interface{}, prefix string) {
	v := reflect.ValueOf(cfg).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("mapstructure")
		if tag == "" {
			tag = field.Name
		}

		tag = strings.ToUpper(tag)
		tag = strings.ReplaceAll(tag, ".", "_")

		if field.Type.Kind() == reflect.Struct {
			BindEnvs(v.Field(i).Addr().Interface(), prefix+tag+"_")
		} else {
			viper.BindEnv(tag, prefix+tag)
		}
	}
}
