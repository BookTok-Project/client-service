package service

import (
	"client-service/internal/db/pg"
	"context"
)

type repo interface {
	GetSubscribeeIDs(ctx context.Context, subscriberID int64) ([]int64, error)
	GetSubscriberIDs(ctx context.Context, subscribeeID int64) ([]int64, error)
	Subscribe(ctx context.Context, subscriberID, subscribeeID int64) error

	AddFavorite(ctx context.Context, userID, bookID int64) error
	RemoveFavorite(ctx context.Context, userID, bookID int64) error
	ListFavorites(ctx context.Context, userID int64) ([]int64, error)

	AddComment(ctx context.Context, userID, bookID int64, text string) error
	GetCommentsByBookID(ctx context.Context, page, limit, bookID int64) ([]pg.CommentsBook, error)
	GetCommentsByUserID(ctx context.Context, page, limit, userID int64) ([]pg.CommentsBook, error)
	GetCountCommentsByBookID(ctx context.Context, bookID int64) (int64, error)
	GetCountCommentsByUserID(ctx context.Context, userID int64) (int64, error)

	AddComplaint(ctx context.Context, userID, bookID int64, text string) error
	GetComplaints(ctx context.Context) ([]pg.Complaint, error)

	AddLike(ctx context.Context, userID, bookID int64) error
	RemoveLike(ctx context.Context, userID, bookID int64) error
	ListLike(ctx context.Context, userID int64) ([]int64, error)
}
