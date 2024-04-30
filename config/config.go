package config

import (
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		HTTP  HTTP  `yaml:"http" env-prefix:"http"`
		Auth  Auth  `yaml:"auth" env-prefix:"auth"`
		MONGO MONGO `yaml:"mongo"`
		JWT   JWT   `yaml:"jwt"`
	}

	HTTP struct {
		Port           string        `env-required:"true" yaml:"port"`
		ReadTimeout    time.Duration `env-required:"true" yaml:"read_timeout"`
		WriteTimeout   time.Duration `env-required:"true" yaml:"write_timeout"`
		IdleTimeout    time.Duration `env-required:"true" yaml:"idle_timeout"`
		MaxHeaderBytes int           `env-required:"true" yaml:"max_header_bytes"`
	}

	MONGO struct {
		URL      string `env:"MONGO_URL" env-default:"mongodb://localhost:27017"`
		Database string `yaml:"database" env-required:"true"`
	}

	Auth struct {
		PasswordCostBcrypt int           `env-required:"true" yaml:"password_cost_bcrypt"`
		AccessTokenTTL     time.Duration `env-required:"true" yaml:"access_token_ttl"`
		RefreshTokenTTL    time.Duration `env-required:"true" yaml:"refresh_token_ttl"`
	}

	JWT struct {
		SigningKey       string `env:"JWT_SIGNING_KEY" env-default:"wdkadwadwakpklrbjb"`
		RefreshTokenCost int    `yaml:"refresh_token_cost_bcrypt"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	configPath := "./config/config.yml"
	path, exists := os.LookupEnv("CONFIG_PATH")
	if exists {
		configPath = path
	}

	err := cleanenv.ReadConfig(configPath, cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}

func MustConfig() *Config {
	cfg, err := NewConfig()
	if err != nil {
		panic(err)
	}
	return cfg
}
