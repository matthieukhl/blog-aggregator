-- name: GetFeeds :many
SELECT f.name, f.url, u.name AS username FROM feeds f INNER JOIN users u ON u.id = f.user_id;