-- name: CreateFeed :one
INSERT into feeds (id, name,url,created_at,updated_at,user_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;


-- name: GetFeedByID :one
SELECT * FROM feeds where id = $1;


-- name: GetNextFeedToFetch :many
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST LIMIT $1;

-- name: MarkFieldAsFetched :one
UPDATE feeds SET last_fetched_at = NOW() , updated_at = NOW() WHERE id = $1
RETURNING *;