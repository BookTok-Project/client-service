package transport

import (
	"client-service/internal/domain"
	"context"

	"github.com/gofiber/fiber/v2"
)

type requestReader interface {
	ReadAndValidateFiberBody(fiberCtx *fiber.Ctx, request any) error
}

type service interface {
	Subscribe(ctx context.Context, subscriberID, subscribeeID int64) error
	GetSubscriberIDs(ctx context.Context, subscribeeID int64) ([]int64, error)
	GetSubscribeeIDs(ctx context.Context, subscriberID int64) ([]int64, error)

	AddFavorite(ctx context.Context, userID, bookID int64) error
	RemoveFavorite(ctx context.Context, userID, bookID int64) error
	ListFavorites(ctx context.Context, userID int64) ([]int64, error)

	AddComment(ctx context.Context, userID, bookID int64, text string) error
	GetCommentsByBookID(ctx context.Context, page, limit, bookID int64) ([]domain.Comment, int64, error)
	GetCommentsByUserID(ctx context.Context, page, limit, userID int64) ([]domain.Comment, int64, error)
}
