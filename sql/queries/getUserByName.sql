-- name: GetUser :one
SELECT id, name, created_at, updated_at FROM users WHERE name = $1;