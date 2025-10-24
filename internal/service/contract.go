package service

import "context"

type repo interface {
	GetSubscribeeIDs(ctx context.Context, subscriberID int64) ([]int64, error)
	GetSubscriberIDs(ctx context.Context, subscribeeID int64) ([]int64, error)
	Subscribe(ctx context.Context, subscriberID, subscribeeID int64) error

	AddFavorite(ctx context.Context, userID, bookID int64) error
	RemoveFavorite(ctx context.Context, userID, bookID int64) error
	ListFavorites(ctx context.Context, userID int64) ([]int64, error)
}
