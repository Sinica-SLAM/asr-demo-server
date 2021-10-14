package handler

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/go-chi/chi/v5"
)

var asrKindPortMap = map[string]int{
	"formospeech_me_1":  8888,
	"tailo_0630":        8889,
	"tailo_0630_taibun": 8890,
	"kenkone":           8891,
	"vgh":               8892,
}

func RegisterWebsocketHandler(router *chi.Mux) {

	router.Route("/websocket", func(websocketRouter chi.Router) {
		websocketRouter.HandleFunc("/{asrKind}/{operation}", func(w http.ResponseWriter, r *http.Request) {
			asrKind := chi.URLParam(r, "asrKind")
			port, ok := asrKindPortMap[asrKind]
			if !ok {
				w.WriteHeader(404)
				return
			}

			operation := chi.URLParam(r, "operation")
			if operation != "speech" && operation != "status" {
				w.WriteHeader(400)
				return
			}

			url, _ := url.Parse(fmt.Sprintf("http://localhost:%d/client/ws/%s", port, operation))

			proxy := httputil.ReverseProxy{Director: func(r *http.Request) {
				r.URL.Scheme = url.Scheme
				r.URL.Host = url.Host
				r.URL.Path = url.Path
				r.Host = url.Host
			}}
			proxy.ServeHTTP(w, r)
		})
	})
}
