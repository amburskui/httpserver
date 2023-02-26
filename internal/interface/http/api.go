package http

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type API struct {
	log    *logrus.Logger
	server *http.Server
}

func New(log *logrus.Logger) *API {
	mux := registerRoute(log)

	return &API{
		log: log,
		server: &http.Server{
			Addr:              ":8000",
			Handler:           mux,
			ReadHeaderTimeout: 3 * time.Second,
		},
	}
}

func (a *API) ListenAndServe(addrAndPort string) error {
	a.log.WithField("port", addrAndPort).Info("listen and server")

	a.server.Addr = addrAndPort

	return a.server.ListenAndServe()
}
