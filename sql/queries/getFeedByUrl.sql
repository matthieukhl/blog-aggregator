-- name: GetFeedByURL :one
SELECT id, name, url, user_id, created_at, updated_at FROM feeds WHERE url = $1;