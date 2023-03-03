-- name: CreateShop :one
INSERT INTO shops (
  name,
  link
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetShop :one
SELECT * FROM shops
WHERE id = $1 LIMIT 1;

-- name: ListShops :many
SELECT * FROM shops
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateShop :one
UPDATE shops
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteShop :exec
DELETE FROM shops
WHERE id = $1;
