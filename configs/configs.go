package configs

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v7"
	"gopkg.in/yaml.v3"
)

type Database struct {
	Host     string `yaml:"host" env:"DATABASE_HOST,required"`
	Port     int    `yaml:"port" env:"DATABASE_PORT,required"`
	Database string `yaml:"database" env:"DATABASE_DATABASE,required"`
	Username string `yaml:"username" env:"DATABASE_USERNAME,required"`
	Password string `yaml:"password" env:"DATABASE_PASSWORD,required"`
}

func (d *Database) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		d.Host, d.Port, d.Username, d.Database, d.Password,
	)
}

type Config struct {
	Listen   string   `yaml:"listen" env:"LISTEN_ADDR"`
	Database Database `yaml:"database"`
}

func Parse(path string) (*Config, error) {
	var config Config

	if path != "" {
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}

		if err := yaml.Unmarshal(data, &config); err != nil {
			return nil, err
		}
	} else {
		if err := env.Parse(&config); err != nil {
			return nil, err
		}
	}

	if config.Listen == "" {
		config.Listen = ":8000"
	}

	return &config, nil
}
