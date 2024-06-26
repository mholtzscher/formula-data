-- name: GetRace :one
SELECT * FROM race
WHERE id = $1 LIMIT 1;
