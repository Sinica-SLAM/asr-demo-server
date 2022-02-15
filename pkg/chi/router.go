package chi

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"go.uber.org/fx"
)

func NewRouter(lifecycle fx.Lifecycle) (*chi.Mux, error) {

	r := chi.NewRouter()
	r.Use(middleware.Logger, middleware.Recoverer, cors.AllowAll().Handler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					panic(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})

	return r, nil
}
