package config

import (
	"context"
	"os"

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
		Service Service `mapstructure:",squash"`
		DB      DB      `mapstructure:",squash"`
		JWT     JWT     `mapstructure:",squash"`
	}

	Service struct {
		Port string `mapstructure:"SERVICE_PORT"`
	}
)

func InitConfig(ctx context.Context) *Configuration {
	var cfg Configuration

	viper.SetConfigType("env")
	envFile := os.Getenv("ENV_FILE")
	if envFile == "" {
		envFile = ".env"
	}

	_, err := os.Stat(envFile)
	if !os.IsNotExist(err) {
		viper.SetConfigFile(envFile)

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("failed to read config:%v", err)
		}
	}

	viper.AutomaticEnv()

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("failed to bind config:%v", err)
	}

	validate := validator.New()
	if err := validate.Struct(cfg); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			log.Errorf("invalid config:%v\n", err)
		}
		log.Fatal("failed to load config")
	}

	return &cfg
}
