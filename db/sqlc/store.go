package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/shopspring/decimal"
)

type Store interface {
	Querier
	ProductTx(ctx context.Context, arg ProductTxParams) (ProductTxResult, error)
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

type ProductTxParams struct {
	Name               string          `json:"name"`
	Price              decimal.Decimal `json:"price"`
	ProductsCategories []int32         `json:"products_categories"`
	Links              []string        `json:"links"`
}

type ProductTxResult struct {
}

func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

func (store *SQLStore) ProductTx(ctx context.Context, arg ProductTxParams) (ProductTxResult, error) {

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		product, err := q.CreateProduct(ctx, CreateProductParams{
			Name:  arg.Name,
			Links: arg.Links,
			Price: sql.NullString{String: arg.Price.String()},
		})
		if err != nil {
			return err
		}

		for _, category := range arg.ProductsCategories {
			_, err = q.InsertNewProductsCategoriesRelationship(ctx, InsertNewProductsCategoriesRelationshipParams{
				ProductCategoryID: sql.NullInt32{Int32: category, Valid: true},
				ProductID:         sql.NullInt32{Int32: product.ID, Valid: true},
			})
			if err != nil {
				return err
			}
		}

		return nil
	})

	return ProductTxResult{}, err
}
