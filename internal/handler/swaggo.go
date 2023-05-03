package handler

import (
	"net/http"

	_ "asr-demo-server/docs"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title ASR DEMO API
// @version 0.1
// @description This is a API for ASR DEMO.

// @contact.name API Support
// @contact.email wayne900619@gmail.com

// @schemes https
// @host asrvm.iis.sinica.edu.tw:8080
// @BasePath /api

func RegisterSwaggoHandler(router *chi.Mux) {
	router.Get("/demo/swagger", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/demo/swagger/index.html", http.StatusMovedPermanently)
	})
	router.Get("/demo/swagger/*", httpSwagger.WrapHandler)

}
