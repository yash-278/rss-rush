-- +goose Up
CREATE TABLE feeds_follows (
    id UUID PRIMARY KEY NOT NULL,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    UNIQUE (feed_id, user_id)
);

-- +goose Down
DROP TABLE feeds_follows;

