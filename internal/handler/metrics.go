package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RegisterMetricsHandler(router *chi.Mux) {
	router.Handle("/metrics", promhttp.Handler())
}
