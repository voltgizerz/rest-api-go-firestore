package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/voltgizerz/rest-api-go-firestore/pkg/logger"
)

type (
	Config struct {
		App       `yaml:"app"`
		HTTP      `yaml:"http"`
		Database  `yaml:"database"`
		Firestore `yaml:"firestore"`
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

	Firestore struct {
		Type                    string `env-required:"true" yaml:"type"`
		ProjectID               string `env-required:"true" yaml:"project_id"`
		PrivateKeyID            string `env-required:"true" yaml:"private_key_id"`
		PrivateKey              string `env-required:"true" yaml:"private_key"`
		ClientEmail             string `env-required:"true" yaml:"client_email"`
		ClientID                string `env-required:"true" yaml:"client_id"`
		AuthURI                 string `env-required:"true" yaml:"auth_uri"`
		TokenURI                string `env-required:"true" yaml:"token_uri"`
		AuthProviderX509CertURL string `env-required:"true" yaml:"auth_provider_x509_cert_url"`
		ClientX509CertURL       string `env-required:"true" yaml:"client_x509_cert_url"`
		UniverseDomain          string `env-required:"true" yaml:"universe_domain"`
	}

	Auth struct {
		JWTSecretKey string `env-required:"true" env:"JWT_SECRET_KEY"`
	}
)

// NewConfig returns app config.
func NewConfig() *Config {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config-dev.yml", cfg)
	if err != nil {
		logger.Log.Fatalln(err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		logger.Log.Fatalln(err)
	}

	return cfg
}