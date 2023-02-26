package http

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func registerRoute(log *logrus.Logger) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/health/", withLogging(log, http.HandlerFunc(healthHandler)))

	return mux
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
}
