-- name: CreateUser :one
INSERT INTO users(id, name,created_at,updated_at)
values ($1,$2,$3,$4)
RETURNING *;