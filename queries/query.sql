-- name: SelectUlidFromId :one
SELECT p.*
FROM product_code_type pct
JOIN product p
    ON pct.id = p.product_code_type_id
WHERE pct.display_name = $1 
    AND p.product_code = $2;
