package api

import "github.com/go-chi/chi/v5"

func Setup(app *api.App) *chi.Mux {
	router := chi.NewRouter()

	router.Route("/products", func(r chi.Router) {
		r.Get("/")
	})
}
