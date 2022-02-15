package handler

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/go-chi/chi/v5"
	httpmiddleware "github.com/slok/go-http-metrics/middleware"
	"github.com/slok/go-http-metrics/middleware/std"
)

var asrKindPortMap = map[string]int{
	"sa_me_2.0":         8888,
	"sa_te_1.0":         8889,
	"sa_me_2.0+kenkone": 8890,
	"sa_me_2.0+vgh":     8891,
}

func RegisterWebsocketHandler(router *chi.Mux, mdlw httpmiddleware.Middleware) {

	router.Route("/websocket", func(websocketRouter chi.Router) {
		websocketRouter.Use(std.HandlerProvider("/websocket/:asrKind/:operation", mdlw))
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
