package main

import (
	"asr-demo-recognize/internal/handler"
	pkg_chi "asr-demo-recognize/pkg/chi"

	"go.uber.org/fx"
)

func main() {
	fx.New(fx.Provide(pkg_chi.NewRouter), fx.Invoke(handler.RegisterStaticHandler, handler.RegisterApiHandler, handler.RegisterSwaggoHandler)).Run()
}
