-- name: GetDriver :one
SELECT * FROM driver
WHERE id = $1 LIMIT 1;
