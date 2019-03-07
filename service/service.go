package service

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-chi/chi"
)

// New sets up a router
func New() chi.Router {
	r := chi.NewRouter()
	r.Use(useLogger)
	r.Get("/metrics", promhttp.Handler().ServeHTTP)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	})
	return r
}
