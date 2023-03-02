package api

import (
	"net/http"

	"github.com/danilluk1/test-task-6/internal/app/api"
	"github.com/danilluk1/test-task-6/internal/app/api/handlers/products"
	"github.com/danilluk1/test-task-6/internal/app/api/middlewares"
	"github.com/go-chi/chi/v5"
)

func Setup(app *api.App) *chi.Mux {
	router := chi.NewRouter()

	router.Route("/products", func(r chi.Router) {
		r.Get("/", products.GetProducts(app))
		r.Get("/{id}", products.GetProduct(app))
		r.With(func(handler http.Handler) http.Handler {
			return middlewares.ValidateAndAttachBody(handler, app, products.CreateProductReq{})
		}).Post("/", products.CreateProduct(app))
		r.Delete("/{id}", products.RemoveProduct(app))

	})
}
