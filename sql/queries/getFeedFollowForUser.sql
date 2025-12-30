-- name: GetFeedFollowsForUser :many
SELECT 
    f.name AS feed_name,
    u.name AS user_name
FROM 
    users u
    INNER JOIN feed_follows ff ON ff.user_id = u.id 
    INNER JOIN feeds f ON f.id = ff.feed_id
WHERE 
    u.name = $1;