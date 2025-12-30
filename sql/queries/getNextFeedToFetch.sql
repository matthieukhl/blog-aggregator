-- name: GetNextFeedToFetch :one
SELECT
    id, name, url
FROM 
    feeds 
ORDER BY 
    last_fetched_at ASC NULLS FIRST
LIMIT 1;