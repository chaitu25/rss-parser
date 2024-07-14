-- name: CreatePosts :many
INSERT into posts(id, created_at,updated_at,tile,description,published_at,url,feed_id) 
VALUES ($1, $2, $3, $4, $5, $6,$7,$8) 
RETURNING *;


-- name: GetPostsForUser :many
SELECT p.*
FROM posts p JOIN feed_followers f 
ON p.feed_id=f.feed_id
WHERE f.user_id = $1
ORDER BY p.published_at DESC 
LIMIT $2;