package products

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
	"github.com/shopspring/decimal"
)

type CreateProductReq struct {
	Name               string          `json:"name" validate:"required,min=5,max=200"`
	Price              decimal.Decimal `json:"price" validate:"required"`
	ProductsCategories []int32         `json:"products_categories"`
	Links              []string        `json:"links" validate:"required"`
}

func CreateProduct(app *api.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		dto := r.Context().Value("body").(*CreateProductReq)
		if dto.Price.LessThanOrEqual(decimal.NewFromInt(0)) {
			response := api_errors.CreateBadRequestError([]string{"price must be higher or equals to 0"})
			w.WriteHeader(http.StatusBadRequest)
			w.Write(response)
			return
		}
		_, err := app.Store.ProductTx(r.Context(), db.ProductTxParams{
			Name:               dto.Name,
			Price:              dto.Price,
			ProductsCategories: dto.ProductsCategories,
			Links:              dto.Links,
		})
		if err != nil {
			app.Logger.Error(err)
			if pqErr, ok := err.(*pq.Error); ok {
				switch pqErr.Code.Name() {
				case "foreign_key_violation", "unique_violation", "not_null_violation":
					response := api_errors.CreateBadRequestError([]string{"wrong products categories"})
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

func GetProduct(app *api.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		productId, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response := api_errors.CreateBadRequestError([]string{"wrong id"})
			w.WriteHeader(http.StatusBadRequest)
			w.Write(response)
			return
		}

		product, err := app.Store.GetProduct(r.Context(), int32(productId))
		if err != nil {
			app.Logger.Error(err)
			if err == sql.ErrNoRows {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		data, err := json.Marshal(product)
		if err != nil {
			app.Logger.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Write(data)
	}
}

func GetProducts(app *api.App) func(w http.ResponseWriter, r *http.Request) {
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

		products, err := app.Store.ListProducts(r.Context(), db.ListProductsParams{
			Limit:  int32(limit),
			Offset: int32(offset),
		})
		if err != nil {
			app.Logger.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		data, err := json.Marshal(products)
		if err != nil {
			app.Logger.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Write(data)
	}
}

func RemoveProduct(app *api.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		productId, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			app.Logger.Error(err)
			response := api_errors.CreateBadRequestError([]string{"Id must be a number"})
			w.WriteHeader(http.StatusBadRequest)
			w.Write(response)
			return
		}

		_, err = app.Store.GetProduct(r.Context(), int32(productId))
		if err != nil {
			app.Logger.Error(err)
			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err = app.Store.DeleteProductsCategoriesRelationshipByProductId(r.Context(), int32(productId))
		if err != nil {
			app.Logger.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err = app.Store.DeleteProduct(r.Context(), int32(productId))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

type CreateCategoryReq struct {
	Name   string `json:"name" validate:"required,min=5,max=200"`
	Link   string `json:"link" validate:"required"`
	ShopId int32  `json:"shop_id" validate:"required"`
}

func CreateCategory(app *api.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func RemoveCategory(app *api.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func GetCategories(app *api.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
