package delivery

import "github.com/gofiber/fiber/v2"

type transport interface {
	HelloWorld(fiberCtx *fiber.Ctx) error
	Subscribe(fiberCtx *fiber.Ctx) error
	GetSubscriberIDs(fiberCtx *fiber.Ctx) error
	GetSubscribeeIDs(fiberCtx *fiber.Ctx) error

	AddFavorite(fiberCtx *fiber.Ctx) error
	RemoveFavorite(fiberCtx *fiber.Ctx) error
	ListFavorites(fiberCtx *fiber.Ctx) error

	AddComment(fiberCtx *fiber.Ctx) error
	GetCommentsByBookID(fiberCtx *fiber.Ctx) error
	GetCommentsByUserID(fiberCtx *fiber.Ctx) error
}
