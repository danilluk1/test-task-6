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

-- name: GetShopsCategories :many
SELECT * FROM shops_categories
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateShopCategory :one
UPDATE shops_categories
SET name = $2
WHERE id = $1
RETURNING *;

-- name: InsertNewShopCategoriesRelationship :one
INSERT INTO shops_shops_categories(
  shop_category_id,
  shop_id
) VALUES (
  $1, $2
) RETURNING *;

-- name: DeleteShopsCategoriesRelationshipByShopId :exec
DELETE FROM shops_shops_categories WHERE shop_id = $1;

-- name: DeleteShopsCategoriesRelationshipByShopsCategoryId :exec
DELETE FROM shops_shops_categories WHERE shop_category_id = $1;

-- name: DeleteShopsCategory :exec
DELETE FROM shops_categories WHERE id = $1;