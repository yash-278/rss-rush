-- +goose Up
-- Add a column to the users table to store the API key and assign default values to existing users.
ALTER TABLE users ADD COLUMN apikey VARCHAR(64) UNIQUE NOT NULL DEFAULT encode(sha256(random()::text::bytea), 'hex');

-- +goose Down
ALTER TABLE users DROP COLUMN apikey;