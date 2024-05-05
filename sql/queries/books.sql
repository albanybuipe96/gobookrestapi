-- name: CreateBook :one
INSERT INTO books (id, title, author, description, pages, user_id, created_at, updated_at)
VALUES($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetByTitle :one
SELECT * FROM books WHERE title = $1;

-- name: GetById :one
SELECT * FROM books WHERE id = $1;

-- name: GetBooks :many
SELECT * FROM books;

-- name: GetBooksByAuthor :many
SELECT * FROM books WHERE author = $1;

-- name: GetBooksByUser :many
SELECT * FROM books WHERE user_id = $1;