package products

import (
	"net/http"
	"strconv"

	"github.com/danilluk1/test-task-6/internal/app/api"
	"github.com/go-chi/chi/v5"
)

func GetProduct(app *api.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		productId, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			a
		}
	}
}

func GetProducts(app *api.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := app.Store.
	}
}
