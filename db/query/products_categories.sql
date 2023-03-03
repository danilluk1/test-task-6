-- name: InsertNewProductsCategoriesRelationship :one
INSERT INTO products_products_categories(
  product_category_id,
  product_id
) VALUES (
  $1, $2
) RETURNING *;

-- name: CreateProductsCategory :one
INSERT INTO products_categories (
  name,
  link,
  shop_id
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetProductsCategories :many
SELECT * FROM products_categories
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteProductsCategoriesRelationshipByProductId :exec
DELETE FROM products_products_categories WHERE product_id = $1;

-- name: DeleteProductsCategoriesRelationshipByProductCategoryId :exec
DELETE FROM products_products_categories WHERE product_category_id = $1;

-- name: DeleteProductsCategory :exec
DELETE FROM products_categories WHERE id = $1;