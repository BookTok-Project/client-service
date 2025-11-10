package transport

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"fmt"
)

func (t *Transport) GetComplatints(fiber* Ctx) error {
	// тут тупо делаем селект по БД	
}

func (t *Transport) AddComplaint(fiberCtx *fiber.Ctx) error {
	userIDStr := fiberCtx.Params("user_id")
	if userID == "" {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Error: "user_id is required",
		})
	}
	
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil || userID < 1 {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Error: "invalid user_id parameter",
		})
	}
	
	bookIDStr := fiberCtx.Params("book_id")
	if bookID == "" {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Error: "book_id is required",
		})
	}
	
	bookID, err := strconv.ParseInt(bookIDStr, 10, 64)
	if err != nil || bookID < 1 {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Error: "invalid book_id parameter",
		})
	}
	
	text := fiberCtx.Params("text")
	if text == "" {
		fiberCtx.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Error: "text is required"
		})
	}

	complaint, err := t.service.AddComplaint(fiberCtx.Context(), userID, bookID, text)
	response := ConvertToAddComplaintsResponse(complaint)
	return fiberCtx.Status(fiber.StatusCreated).JSON(response)
}
