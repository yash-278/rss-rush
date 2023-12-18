-- name: GetFeedByIdAndUserId :one
SELECT * FROM feeds_follows WHERE feed_id = $1 AND user_id = $2;

-- name: CreateFeedFollow :one
INSERT INTO feeds_follows (id, feed_id, user_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;