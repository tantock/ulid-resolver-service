-- name: SelectUlidFromId :one
SELECT p.*
FROM product_code_type pct
JOIN product p
    ON pct.id = p.product_code_type_id
WHERE pct.display_name = $1 
    AND p.product_code = $2;

-- name: InsertProduct :one
INSERT INTO product (id, product_code, product_code_type_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: SelectProductCodeTypeByName :one
SELECT *
FROM product_code_type pct
WHERE pct.display_name = $1;