-- name: CreatePost :one
INSERT INTO posts (
    id,
    title,
    description,
    url,
    feed_id,
    published_at,
    created_at,
    updated_at
) 
VALUES($1,$2,$3,$4,$5,$6,$7,$8) 
RETURNING *;


-- name: GetPostsForUser :many
SELECT posts.* from posts
JOIN feed_followers ON posts.feed_id = feed_followers.feed_id
WHERE feed_followers.user_id = $1
ORDER BY posts.published_at DESC
LIMIT $2;