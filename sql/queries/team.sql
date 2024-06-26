-- name: GetTeam :one
SELECT * FROM team
WHERE id = $1 LIMIT 1;

-- name: ListTeams :many
SELECT * FROM team ORDER BY id;
