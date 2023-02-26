package main

import (
	"flag"
	"fmt"

	"github.com/sirupsen/logrus"

	httpAPI "github.com/amburskui/httpserver/internal/interface/http"
)

func main() {
	port := flag.Int("p", 8000, "PORT")
	flag.Parse()

	log := logrus.New()

	api := httpAPI.New(log)
	api.ListenAndServe(fmt.Sprintf(":%d", *port))
}
