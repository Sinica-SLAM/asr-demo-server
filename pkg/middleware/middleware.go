package middleware

import (
	httpmetrics "github.com/slok/go-http-metrics/metrics/prometheus"
	httpmiddleware "github.com/slok/go-http-metrics/middleware"
	"go.uber.org/fx"
)

func NewHttpMetricsMiddleware(lifecycle fx.Lifecycle) httpmiddleware.Middleware {
	return httpmiddleware.New(httpmiddleware.Config{Recorder: httpmetrics.NewRecorder(httpmetrics.Config{})})
}
