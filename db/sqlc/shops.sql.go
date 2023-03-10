// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: shops.sql

package db

import (
	"context"
)

const createShop = `-- name: CreateShop :one
INSERT INTO shops (
  name,
  link
) VALUES (
  $1, $2
) RETURNING id, name, link, created_at
`

type CreateShopParams struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

func (q *Queries) CreateShop(ctx context.Context, arg CreateShopParams) (Shop, error) {
	row := q.db.QueryRowContext(ctx, createShop, arg.Name, arg.Link)
	var i Shop
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Link,
		&i.CreatedAt,
	)
	return i, err
}

const deleteShop = `-- name: DeleteShop :exec
DELETE FROM shops
WHERE id = $1
`

func (q *Queries) DeleteShop(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteShop, id)
	return err
}

const getShop = `-- name: GetShop :one
SELECT id, name, link, created_at FROM shops
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetShop(ctx context.Context, id int32) (Shop, error) {
	row := q.db.QueryRowContext(ctx, getShop, id)
	var i Shop
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Link,
		&i.CreatedAt,
	)
	return i, err
}

const listShops = `-- name: ListShops :many
SELECT id, name, link, created_at FROM shops
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListShopsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListShops(ctx context.Context, arg ListShopsParams) ([]Shop, error) {
	rows, err := q.db.QueryContext(ctx, listShops, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Shop{}
	for rows.Next() {
		var i Shop
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Link,
			&i.CreatedAt,
		); err != nil {
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

const updateShop = `-- name: UpdateShop :one
UPDATE shops
SET name = $2
WHERE id = $1
RETURNING id, name, link, created_at
`

type UpdateShopParams struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) UpdateShop(ctx context.Context, arg UpdateShopParams) (Shop, error) {
	row := q.db.QueryRowContext(ctx, updateShop, arg.ID, arg.Name)
	var i Shop
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Link,
		&i.CreatedAt,
	)
	return i, err
}
