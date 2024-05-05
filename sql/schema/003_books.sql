-- +goose Up

CREATE TABLE books(
    id UUID PRIMARY KEY,
    title TEXT UNIQUE NOT NULL,
    author TEXT NOT NULL,
    description TEXT UNIQUE NOT NULL,
    pages INTEGER NOT NULL DEFAULT(0),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE books;