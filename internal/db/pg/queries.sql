-- name: GetSubscriberIDs :many
SELECT subscriber_id
FROM subscriptions
WHERE subscribee_id = $1;

-- name: GetSubscribeeIDs :many
SELECT subscribee_id
FROM subscriptions
WHERE subscriber_id = $1;

-- name: Subscribe :exec
INSERT INTO subscriptions (subscriber_id, subscribee_id)
VALUES ($1, $2);
