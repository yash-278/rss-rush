-- +goose Up
CREATE TABLE feeds (
    id UUID PRIMARY KEY NOT NULL,
    url VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    UNIQUE (url)
);

-- +goose Down
DROP TABLE feeds;
