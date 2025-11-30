package transport

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (t *Transport) AddLike(fiberCtx *fiber.Ctx) error {
	var req likeRequest
	if err := t.requestReader.ReadAndValidateFiberBody(fiberCtx, &req); err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := t.service.AddLike(fiberCtx.Context(), req.UserID, req.BookID); err != nil {
		return fiberCtx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return fiberCtx.SendStatus(fiber.StatusCreated)
}

func (t *Transport) RemoveLike(fiberCtx *fiber.Ctx) error {
	var req likeRequest
	if err := t.requestReader.ReadAndValidateFiberBody(fiberCtx, &req); err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := t.service.RemoveLike(fiberCtx.Context(), req.UserID, req.BookID); err != nil {
		return fiberCtx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return fiberCtx.SendStatus(fiber.StatusNoContent)
}

func (t *Transport) ListLike(fiberCtx *fiber.Ctx) error {
	userIDParam := fiberCtx.Params("user_id")
	if userIDParam == "" {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user_id required"})
	}

	userID, err := strconv.ParseInt(userIDParam, 10, 64)
	if err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user_id"})
	}

	ids, err := t.service.ListLike(fiberCtx.Context(), userID)
	if err != nil {
		return fiberCtx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(ids)
}
