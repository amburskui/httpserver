package main

import (
	"errors"
	"flag"
	"net/http"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/amburskui/httpserver/configs"
	"github.com/amburskui/httpserver/internal/application/userservice"
	"github.com/amburskui/httpserver/internal/domain"
	"github.com/amburskui/httpserver/internal/infrastructure/persistence"
	httpAPI "github.com/amburskui/httpserver/internal/interface/http"
)

func main() {
	opts := struct {
		configPath string
	}{}

	flag.StringVar(&opts.configPath, "c", "", "Configuration FILE")
	flag.Parse()

	log := logrus.New()

	var config configs.Config
	if err := configs.Parse(opts.configPath, &config); err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open(postgres.New(postgres.Config{DSN: config.Database.DSN()}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db = db.Debug()

	storage := persistence.NewStorage(db)
	userService := userservice.New(storage)

	db.AutoMigrate(new(domain.User))

	api := httpAPI.New(log, userService)
	if err := api.ListenAndServe(config.Listen); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return
		}

		log.Fatal(err)
	}
}
