package transport

import (
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
}
