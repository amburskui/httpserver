package main

import (
	"flag"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/amburskui/httpserver/configs"
	"github.com/amburskui/httpserver/internal/domain"
)

func main() {
	opts := struct {
		configPath string
	}{}

	flag.StringVar(&opts.configPath, "c", "configs/dbmigrate.yml", "Configuration FILE")
	flag.Parse()

	log := logrus.New()

	var database configs.Database
	if err := configs.Parse(opts.configPath, &database); err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open(postgres.New(postgres.Config{DSN: database.DSN()}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db = db.Debug()
	db.AutoMigrate(new(domain.User))
}
