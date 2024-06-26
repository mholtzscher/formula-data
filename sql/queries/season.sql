-- name: GetSeason :one
SELECT * FROM season
WHERE id = $1 LIMIT 1;
