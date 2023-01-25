package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func withLogging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.WithFields(logrus.Fields{
			"uri":    r.RequestURI,
			"method": r.Method,
		}).Info("request")

		h.ServeHTTP(w, r)
	})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
}

func main() {
	port := flag.Int("p", 8000, "PORT")
	flag.Parse()

	mux := http.NewServeMux()
	mux.Handle("/health/", withLogging(http.HandlerFunc(healthHandler)))

	logrus.WithField("port", *port).Info("listen and server")

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", *port),
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           mux,
	}

	logrus.Error(server.ListenAndServe())
}
