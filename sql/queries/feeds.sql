-- name: CreateFeed :one
INSERT into feeds (id, name,url,created_at,updated_at,user_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;