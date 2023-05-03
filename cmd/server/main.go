package main

import (
	"asr-demo-server/internal/handler"
	pkg_chi "asr-demo-server/pkg/chi"
	"asr-demo-server/pkg/youtube"

	"go.uber.org/fx"
)

func main() {
	fx.New(fx.Provide(pkg_chi.NewRouter, youtube.NewService), fx.Invoke(handler.RegisterStaticHandler, handler.RegisterApiHandler, handler.RegisterWebsocketHandler, handler.RegisterSwaggoHandler)).Run()
}
