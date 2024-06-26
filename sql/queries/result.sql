-- name: GetResult :one
SELECT * FROM result
WHERE id = $1 LIMIT 1;
