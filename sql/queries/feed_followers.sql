-- name: CreateFeedFollower :one
INSERT into feed_followers (id,created_at,updated_at,user_id,feed_id) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetFeedsFollows :many
SELECT * FROM feed_followers WHERE user_id = $1;
