package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
)

const DATE_FORMAT = "2006-01-02 15:04"
const POSTS_PER_PAGE = 10

type BindConfig struct {
	Host string `env:"HOST" envDefault:"127.0.0.1"`
	Port uint   `env:"PORT" envDefault:"8080"`
}

func (c *BindConfig) AsAddress() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

type AppConfig struct {
	Bind           BindConfig `envPrefix:"BIND_"`
	Email          string     `env:"EMAIL,notEmpty"`
	GitHub         string     `env:"GITHUB,notEmpty"`
	Domain         string     `env:"DOMAIN,notEmpty"`
	PosthogApiKey  string     `env:"POSTHOG_API_KEY,notEmpty"`
	PosthogApiHost string     `env:"POSTHOG_API_HOST,notEmpty"`
}

func NewConfig() (*AppConfig, error) {
	conf := AppConfig{}
	if err := env.Parse(&conf); err != nil {
		return nil, fmt.Errorf("config parse: %w", err)
	}

	return &conf, nil
}
