package transport

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type favoriteRequest struct {
	UserID int64 `json:"user_id"`
	BookID int64 `json:"book_id"`
}

func (t *Transport) AddFavorite(fiberCtx *fiber.Ctx) error {
	var req favoriteRequest
	if err := t.requestReader.ReadAndValidateFiberBody(fiberCtx, &req); err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := t.service.AddFavorite(fiberCtx.Context(), req.UserID, req.BookID); err != nil {
		return fiberCtx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return fiberCtx.SendStatus(fiber.StatusCreated)
}

func (t *Transport) RemoveFavorite(fiberCtx *fiber.Ctx) error {
	var req favoriteRequest
	if err := t.requestReader.ReadAndValidateFiberBody(fiberCtx, &req); err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := t.service.RemoveFavorite(fiberCtx.Context(), req.UserID, req.BookID); err != nil {
		return fiberCtx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return fiberCtx.SendStatus(fiber.StatusNoContent)
}

func (t *Transport) ListFavorites(fiberCtx *fiber.Ctx) error {
	userIDParam := fiberCtx.Params("user_id") // например: /favorites/:user_id
	if userIDParam == "" {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user_id required"})
	}

	userID, err := strconv.ParseInt(userIDParam, 10, 64)
	if err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user_id"})
	}

	ids, err := t.service.ListFavorites(fiberCtx.Context(), userID)
	if err != nil {
		return fiberCtx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(ids)
}
