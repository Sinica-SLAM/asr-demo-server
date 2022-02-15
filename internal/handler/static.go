package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterStaticHandler(router *chi.Mux) {
	router.Handle("/*", http.FileServer(http.Dir("./dist")))
}
