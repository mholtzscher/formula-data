-- name: GetTeam :one
SELECT * FROM team
WHERE id = $1 LIMIT 1;
