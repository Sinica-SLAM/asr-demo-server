package main

import (
	"asr-demo-recognize/internal/handler"
	pkg_chi "asr-demo-recognize/pkg/chi"
	"asr-demo-recognize/pkg/youtube"

	"go.uber.org/fx"
)

func main() {
	fx.New(fx.Provide(pkg_chi.NewRouter, youtube.NewService), fx.Invoke(handler.RegisterStaticHandler, handler.RegisterApiHandler, handler.RegisterWebsocketHandler, handler.RegisterSwaggoHandler, handler.RegisterMetricsHandler)).Run()
}
