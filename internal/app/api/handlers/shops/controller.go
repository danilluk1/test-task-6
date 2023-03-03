package shops

import (
	"database/sql"
	"net/http"
	"strconv"

	db "github.com/danilluk1/test-task-6/db/sqlc"
	"github.com/danilluk1/test-task-6/internal/app/api"
	"github.com/danilluk1/test-task-6/internal/app/api/api_errors"
	"github.com/go-chi/chi/v5"
	"github.com/lib/pq"
)

type CreateShopReq struct {
	Name            string  `json:"name" validate:"required,min=5,max=200"`
	Link            string  `json:"link" validate:"required,min=5,max=200"`
	ShopsCategories []int32 `json:"products_categories"`
}

func CreateShop(app *api.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		dto := r.Context().Value("body").(*CreateShopReq)

		_, err := app.Store.ShopTx(r.Context(), db.ShopTxParams{
			Name:            dto.Name,
			Link:            dto.Link,
			ShopsCategories: dto.ShopsCategories,
		})
		if err != nil {
			app.Logger.Error(err)
			if pqErr, ok := err.(*pq.Error); ok {
				switch pqErr.Code.Name() {
				case "foreign_key_violation", "unique_violation", "not_null_violation":
					response := api_errors.CreateBadRequestError([]string{"wrong shop categories"})
					w.WriteHeader(http.StatusBadRequest)
					w.Write(response)
					return
				}
			}
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
}

func RemoveShop(app *api.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		shopId, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			app.Logger.Error(err)
			response := api_errors.CreateBadRequestError([]string{"Id muste be a number"})
			w.WriteHeader(http.StatusBadRequest)
			w.Write(response)
			return
		}

		_, err = app.Store.GetProduct(r.Context(), int32(shopId))
		if err != nil {
			app.Logger.Error(err)
			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err = app.Store.Dee
	}
}

func GetShops(app *api.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func GetShop(app *api.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func GetCategories(app *api.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func CreateCategory(app *api.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func RemoveCategory(app *api.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
