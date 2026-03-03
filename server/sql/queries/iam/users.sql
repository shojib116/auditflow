-- name: CreateUser :one
INSERT INTO users (email, full_name, password_hash, is_verified)
VALUES ($1, $2, $3, FALSE)
RETURNING *;

