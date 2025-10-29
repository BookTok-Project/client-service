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

-- name: AddFavoriteBook :exec
INSERT INTO user_favorite_books (user_id, book_id)
VALUES ($1, $2)
ON CONFLICT (user_id, book_id) DO NOTHING;

-- name: RemoveFavoriteBook :exec
DELETE FROM user_favorite_books
WHERE user_id = $1 AND book_id = $2;

-- name: ListFavoriteBooks :many
SELECT book_id
FROM user_favorite_books
WHERE user_id = $1
ORDER BY created_at DESC;
