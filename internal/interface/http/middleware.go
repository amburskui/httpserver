package http

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func withLogging(log *logrus.Logger, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(logrus.Fields{
			"uri":    r.RequestURI,
			"method": r.Method,
		}).Info("request")

		h.ServeHTTP(w, r)
	})
}
