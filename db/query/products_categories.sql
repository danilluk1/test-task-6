-- name: InsertNewProductsCategoriesRelationship :one
INSERT INTO products_products_categories(
  product_category_id,
  product_id
) VALUES (
  $1, $2
) RETURNING *;

-- name: DeleteProductsCategoriesRelationshipByProductId :exec
DELETE FROM products_products_categories WHERE product_id = $1;

-- name: DeleteProductsCategoriesRelationshipByProductCategoryId :exec
DELETE FROM products_products_categories WHERE product_category_id = $1;