package handler

import (
	"net/http"

	_ "asr-demo-recognize/docs"

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
	router.Get("/api/swagger", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/api/swagger/index.html", http.StatusMovedPermanently)
	})
	router.Get("/api/swagger/*", httpSwagger.WrapHandler)

}
