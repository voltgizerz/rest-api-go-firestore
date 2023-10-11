package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/voltgizerz/rest-api-go-firestore/pkg/logger"
)

type (
	Config struct {
		App      `yaml:"app"`
		HTTP     `yaml:"http"`
		Database `yaml:"database"`
	}

	App struct {
		Name    string `env-required:"true" yaml:"name"`
		Version string `env-required:"true" yaml:"version"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port"`
	}

	Database struct {
		PoolMax int `env-required:"true" yaml:"pool_max"`
	}
)

// NewConfig returns app config.
func NewConfig() *Config {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		logger.Log.Fatalln(err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		logger.Log.Fatalln(err)
	}

	return cfg
}
