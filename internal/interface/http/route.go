package http

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func registerRoute(log *logrus.Logger) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/health/", withLogging(log, http.HandlerFunc(healthHandler)))
	mux.Handle("/", withLogging(log, http.HandlerFunc(otusStudentHandler)))

	return mux
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
}

func otusStudentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	student := r.URL.Query().Get("student")

	if student == "" {
		student = "anonymous"
	}

	json.NewEncoder(w).Encode(map[string]string{"welcome": student})
}
