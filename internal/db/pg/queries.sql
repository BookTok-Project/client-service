-- name: GetSubscribers :many
SELECT reader_id
FROM subscriptions
WHERE writer_id = $1;

-- name: GetSubscriptions :many
SELECT writer_id
FROM subscriptions
WHERE reader_id = $1;
