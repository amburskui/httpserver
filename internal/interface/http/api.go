package http

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/amburskui/httpserver/internal/application/userservice"
)

type API struct {
	log         *logrus.Logger
	server      *http.Server
	userService *userservice.Service
}

func New(log *logrus.Logger, service *userservice.Service) *API {
	mux := registerRoute(log, service)

	return &API{
		log: log,
		server: &http.Server{
			Addr:              ":8000",
			Handler:           mux,
			ReadHeaderTimeout: 3 * time.Second,
		},
		userService: service,
	}
}

func (a *API) ListenAndServe(addrAndPort string) error {
	a.log.WithField("port", addrAndPort).Info("listen and server")

	a.server.Addr = addrAndPort

	return a.server.ListenAndServe()
}
