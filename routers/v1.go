package routers

import (
	"github.com/albanybuipe96/bookrestapi/handlers"
	"github.com/go-chi/chi"
)

func V1Router() *chi.Mux {
	v1Router := chi.NewRouter()
	v1Router.Get("/", handlers.Index)
	v1Router.Get("/error", handlers.Error)

	return v1Router
}
