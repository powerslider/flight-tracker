package configs

import (
	"github.com/joeshaw/envdecode"
	"github.com/pkg/errors"
)

type ServerConfig struct {
	Host string `env:"SERVER_HOST"`
	Port int    `env:"SERVER_PORT"`
}

func NewServerConfig() (*ServerConfig, error) {
	var config ServerConfig

	err := envdecode.Decode(&config)

	if err != nil {
		if err != envdecode.ErrNoTargetFieldsAreSet {
			return nil, errors.WithStack(err)
		}
	}

	return &config, nil
}
