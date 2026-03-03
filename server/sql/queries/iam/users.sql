-- name: CreateUser :one
INSERT INTO users (email, full_name, password_hash, is_verified)
VALUES ($1, $2, $3, FALSE)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

