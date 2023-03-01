package configs

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func (d *Database) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		d.Host, d.Port, d.Username, d.Database, d.Password,
	)
}

type Config struct {
	Listen   string   `yaml:"listen"`
	Database Database `yaml:"database"`
}

func Parse[T Config | Database](configPath string, config *T) error {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	password := os.Getenv("DATABASE_PASSWORD")

	var c any = config

	switch a := c.(type) {
	case *Config:
		a.Database.Password = password
	case *Database:
		a.Password = password
	}

	return yaml.Unmarshal(data, &config)
}
