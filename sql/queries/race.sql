-- name: GetRaceById :one
SELECT * FROM race
WHERE id = ? LIMIT 1;


