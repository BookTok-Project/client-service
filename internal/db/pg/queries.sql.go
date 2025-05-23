// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: queries.sql

package pg

import (
	"context"
)

const getSubscribers = `-- name: GetSubscribers :many
SELECT reader_id
FROM subscriptions
WHERE writer_id = $1
`

func (q *Queries) GetSubscribers(ctx context.Context, writerID int64) ([]int64, error) {
	rows, err := q.db.Query(ctx, getSubscribers, writerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int64
	for rows.Next() {
		var reader_id int64
		if err := rows.Scan(&reader_id); err != nil {
			return nil, err
		}
		items = append(items, reader_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSubscriptions = `-- name: GetSubscriptions :many
SELECT writer_id
FROM subscriptions
WHERE reader_id = $1
`

func (q *Queries) GetSubscriptions(ctx context.Context, readerID int64) ([]int64, error) {
	rows, err := q.db.Query(ctx, getSubscriptions, readerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int64
	for rows.Next() {
		var writer_id int64
		if err := rows.Scan(&writer_id); err != nil {
			return nil, err
		}
		items = append(items, writer_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
