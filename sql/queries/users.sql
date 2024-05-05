-- name: CreateUser :one
INSERT INTO users (id, username, email, created_at, updated_at)
VALUES($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUserByAPIKey :one
SELECT * FROM users WHERE api_key = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1;