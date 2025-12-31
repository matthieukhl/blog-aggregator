-- name: GetPostsForUser :many
SELECT
    p.id,
    p.title,
    p.description,
    p.url,
    p.published_at,
    f.name AS feed_name
FROM
    posts p 
    INNER JOIN feed_follows ff ON ff.feed_id = p.feed_id
    INNER JOIN feeds f ON f.id = p.feed_id
WHERE 
    ff.user_id = $1
ORDER BY 
    p.published_at DESC 
LIMIT $2;
