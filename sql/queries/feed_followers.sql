-- name: CreateFeedFollow :one
INSERT INTO feed_followers (id, feed_id, user_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;


-- name: GetFeedFollows :many
SELECT * FROM feed_followers WHERE user_id = $1;

-- name: DeleteFeedFollow :exec
DELETE  FROM feed_followers WHERE id = $1 AND user_id = $2;