-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, user_id, feed_id, created_at, updated_at)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
) SELECT iff.*, f.name AS feed_name, u.name AS user_name
FROM inserted_feed_follow iff
INNER JOIN feeds f ON f.id = iff.feed_id 
INNER JOIN users u ON u.id = iff.user_id;