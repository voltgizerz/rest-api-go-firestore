package config

import (
	"context"
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/opentracing/opentracing-go"
	"github.com/voltgizerz/rest-api-go-firestore/pkg/env"
	"github.com/voltgizerz/rest-api-go-firestore/pkg/logger"
)

type (
	Config struct {
		App      `yaml:"app"`
		HTTP     `yaml:"http"`
		Database `yaml:"database"`
		Auth
	}

	App struct {
		Name          string `env-required:"true" yaml:"name"`
		Version       string `env-required:"true" yaml:"version"`
		IsUseFakeData bool   `env-required:"true" yaml:"is_use_fake_data"`
	}

	// HTTP -.
	HTTP struct {
		GinMode string `env-required:"true" yaml:"gin_mode"`
		Port    string `env-required:"true" yaml:"port"`
	}

	Database struct {
		PoolMax int `env-required:"true" yaml:"pool_max"`
	}

	Auth struct {
		JWTSecretKey string `env-required:"true" env:"JWT_SECRET_KEY"`
	}
)

// NewConfig returns app config.
func NewConfig(ctx context.Context) *Config {
	span, _ := opentracing.StartSpanFromContext(ctx, "config.NewConfig")
	defer span.Finish()

	cfg := &Config{}

	configPath := fmt.Sprintf("./config/file/config-%s.yml", env.GetENV())
	err := cleanenv.ReadConfig(configPath, cfg)
	if err != nil {
		logger.Log.Fatalf("[NewConfig.ReadConfig] path: %s, err: %v", configPath, err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		logger.Log.Fatalf("[NewConfig.ReadEnv] err: %v", err)
	}

	return cfg
}
