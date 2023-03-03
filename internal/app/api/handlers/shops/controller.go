package shops

import (
	"database/sql"
	"encoding/json"
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
	ShopsCategories []int32 `json:"shops_categories" validate:"required"`
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
			response := api_errors.CreateBadRequestError([]string{"Id must be a number"})
			w.WriteHeader(http.StatusBadRequest)
			w.Write(response)
			return
		}

		_, err = app.Store.GetShop(r.Context(), int32(shopId))
		if err != nil {
			app.Logger.Error(err)
			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err = app.Store.DeleteShopsCategoriesRelationshipByShopId(r.Context(), int32(shopId))
		if err != nil {
			app.Logger.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err = app.Store.DeleteShop(r.Context(), int32(shopId))
		if err != nil {
			app.Logger.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
}

func GetShops(app *api.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var limit uint8 = 10
		var offset uint64 = 0

		query := r.URL.Query()
		limitParam := query.Get("limit")
		if len(limitParam) != 0 {
			newLimit, err := strconv.ParseUint(limitParam, 10, 64)
			if err != nil {
				response := api_errors.CreateBadRequestError([]string{"wrong limit"})
				w.WriteHeader(http.StatusBadRequest)
				w.Write(response)
				return
			}
			if newLimit > 30 {
				response := api_errors.CreateBadRequestError([]string{"limit must be lower than 30"})
				w.WriteHeader(http.StatusBadRequest)
				w.Write(response)
				return
			}
			limit = uint8(newLimit)
		}

		shops, err := app.Store.ListShops(r.Context(), db.ListShopsParams{
			Limit:  int32(limit),
			Offset: int32(offset),
		})
		if err != nil {
			app.Logger.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		data, err := json.Marshal(shops)
		if err != nil {
			app.Logger.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Write(data)
	}
}

func GetShop(app *api.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		shopId, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response := api_errors.CreateBadRequestError([]string{"wrong id"})
			w.WriteHeader(http.StatusBadRequest)
			w.Write(response)
			return
		}

		shop, err := app.Store.GetShop(r.Context(), int32(shopId))
		if err != nil {
			app.Logger.Error(err)
			if err == sql.ErrNoRows {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		data, err := json.Marshal(shop)
		if err != nil {
			app.Logger.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Write(data)
	}
}

func GetCategories(app *api.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var limit uint8 = 10
		var offset uint64 = 0

		query := r.URL.Query()
		limitParam := query.Get("limit")
		if len(limitParam) != 0 {
			newLimit, err := strconv.ParseUint(limitParam, 10, 64)
			if err != nil {
				response := api_errors.CreateBadRequestError([]string{"wrong limit"})
				w.WriteHeader(http.StatusBadRequest)
				w.Write(response)
				return
			}
			if newLimit > 30 {
				response := api_errors.CreateBadRequestError([]string{"limit must be lower than 30"})
				w.WriteHeader(http.StatusBadRequest)
				w.Write(response)
				return
			}
			limit = uint8(newLimit)
		}

		categories, err := app.Store.GetShopsCategories(r.Context(), db.GetShopsCategoriesParams{
			Limit:  int32(limit),
			Offset: int32(offset),
		})
		if err != nil {
			app.Logger.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		data, err := json.Marshal(categories)
		if err != nil {
			app.Logger.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Write(data)
	}
}

type CreateCategoryReq struct {
	Name string `json:"name" validate:"required,min=5,max=200"`
	Link string `json:"link" validate:"required"`
}

func CreateCategory(app *api.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		dto := r.Context().Value("body").(*CreateCategoryReq)

		category, err := app.Store.CreateShopCategory(r.Context(), db.CreateShopCategoryParams{
			Name: dto.Name,
			Link: dto.Link,
		})
		if err != nil {
			app.Logger.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		data, err := json.Marshal(category)
		if err != nil {
			app.Logger.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Write(data)
	}
}

func RemoveCategory(app *api.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		categoryId, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			app.Logger.Error(err)
			response := api_errors.CreateBadRequestError([]string{"Id must be a number"})
			w.WriteHeader(http.StatusBadRequest)
			w.Write(response)
			return
		}

		err = app.Store.DeleteShopsCategoriesRelationshipByShopsCategoryId(r.Context(), int32(categoryId))
		if err != nil {
			app.Logger.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err = app.Store.DeleteShopsCategory(r.Context(), int32(categoryId))
		if err != nil {
			app.Logger.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
}
