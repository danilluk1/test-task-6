// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: shops_categories.sql

package db

import (
	"context"
)

const createShopCategory = `-- name: CreateShopCategory :one
INSERT INTO shops_categories (
  name,
  link
) VALUES (
  $1, $2
) RETURNING id, name, link
`

type CreateShopCategoryParams struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

func (q *Queries) CreateShopCategory(ctx context.Context, arg CreateShopCategoryParams) (ShopsCategory, error) {
	row := q.db.QueryRowContext(ctx, createShopCategory, arg.Name, arg.Link)
	var i ShopsCategory
	err := row.Scan(&i.ID, &i.Name, &i.Link)
	return i, err
}

const getShopCategory = `-- name: GetShopCategory :one
SELECT id, name, link FROM shops_categories
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetShopCategory(ctx context.Context, id int32) (ShopsCategory, error) {
	row := q.db.QueryRowContext(ctx, getShopCategory, id)
	var i ShopsCategory
	err := row.Scan(&i.ID, &i.Name, &i.Link)
	return i, err
}

const listShopsCategories = `-- name: ListShopsCategories :many
SELECT id, name, link FROM shops_categories
ORDER BY shop_category_id
LIMIT $1
OFFSET $2
`

type ListShopsCategoriesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListShopsCategories(ctx context.Context, arg ListShopsCategoriesParams) ([]ShopsCategory, error) {
	rows, err := q.db.QueryContext(ctx, listShopsCategories, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ShopsCategory{}
	for rows.Next() {
		var i ShopsCategory
		if err := rows.Scan(&i.ID, &i.Name, &i.Link); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateShopCategory = `-- name: UpdateShopCategory :one
UPDATE shops_categories
SET name = $2
WHERE id = $1
RETURNING id, name, link
`

type UpdateShopCategoryParams struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) UpdateShopCategory(ctx context.Context, arg UpdateShopCategoryParams) (ShopsCategory, error) {
	row := q.db.QueryRowContext(ctx, updateShopCategory, arg.ID, arg.Name)
	var i ShopsCategory
	err := row.Scan(&i.ID, &i.Name, &i.Link)
	return i, err
}
