-- name: CreateShopCategory :one
INSERT INTO shops_categories (
  name,
  link
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetShopCategory :one
SELECT * FROM shops_categories
WHERE id = $1 LIMIT 1;

-- name: ListShopsCategories :many
SELECT * FROM shops_categories
ORDER BY shop_category_id
LIMIT $1
OFFSET $2;

-- name: UpdateShopCategory :one
UPDATE shops_categories
SET name = $2
WHERE id = $1
RETURNING *;
