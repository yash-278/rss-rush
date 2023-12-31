// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: feed_follows.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFeedFollow = `-- name: CreateFeedFollow :one
INSERT INTO feeds_follows (id, feed_id, user_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, feed_id, user_id, created_at, updated_at
`

type CreateFeedFollowParams struct {
	ID        uuid.UUID
	FeedID    uuid.UUID
	UserID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) CreateFeedFollow(ctx context.Context, arg CreateFeedFollowParams) (FeedsFollow, error) {
	row := q.db.QueryRowContext(ctx, createFeedFollow,
		arg.ID,
		arg.FeedID,
		arg.UserID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i FeedsFollow
	err := row.Scan(
		&i.ID,
		&i.FeedID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteFeedFollow = `-- name: DeleteFeedFollow :exec
DELETE FROM feeds_follows WHERE id = $1
`

func (q *Queries) DeleteFeedFollow(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteFeedFollow, id)
	return err
}

const getFeedByIdAndUserId = `-- name: GetFeedByIdAndUserId :one
SELECT id, feed_id, user_id, created_at, updated_at FROM feeds_follows WHERE feed_id = $1 AND user_id = $2
`

type GetFeedByIdAndUserIdParams struct {
	FeedID uuid.UUID
	UserID uuid.UUID
}

func (q *Queries) GetFeedByIdAndUserId(ctx context.Context, arg GetFeedByIdAndUserIdParams) (FeedsFollow, error) {
	row := q.db.QueryRowContext(ctx, getFeedByIdAndUserId, arg.FeedID, arg.UserID)
	var i FeedsFollow
	err := row.Scan(
		&i.ID,
		&i.FeedID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getFeedFollowsByUserId = `-- name: GetFeedFollowsByUserId :many
SELECT id, feed_id, user_id, created_at, updated_at FROM feeds_follows WHERE user_id = $1
`

func (q *Queries) GetFeedFollowsByUserId(ctx context.Context, userID uuid.UUID) ([]FeedsFollow, error) {
	rows, err := q.db.QueryContext(ctx, getFeedFollowsByUserId, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FeedsFollow
	for rows.Next() {
		var i FeedsFollow
		if err := rows.Scan(
			&i.ID,
			&i.FeedID,
			&i.UserID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
