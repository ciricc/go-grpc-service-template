package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port string `yaml:"port" env:"PORT"`
}

func New(file string) (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadConfig(file, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
