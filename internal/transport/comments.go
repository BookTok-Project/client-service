package transport

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (t *Transport) AddComment(fiberCtx *fiber.Ctx) error {
	var req addCommentRequest
	if err := t.requestReader.ReadAndValidateFiberBody(fiberCtx, &req); err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := t.service.AddComment(fiberCtx.Context(), req.UserID, req.BookID, req.Text)
	if err != nil {
		return fmt.Errorf("AddComment(%d, %d, %s): %w", req.UserID, req.BookID, req.Text, err)
	}

	return fiberCtx.SendStatus(fiber.StatusCreated)
}

func (t *Transport) GetCommentsByBookID(fiberCtx *fiber.Ctx) error {
	pageQuery := fiberCtx.Query("page")
	if pageQuery == "" {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Error: "page parameter is required",
		})
	}

	page, err := strconv.Atoi(pageQuery)
	if err != nil || page < 0 {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Error: "invalid page parameter: must be a non negative integer",
		})
	}

	limitQuery := fiberCtx.Query("limit")
	if limitQuery == "" {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Error: "limit parameter is required",
		})
	}

	limit, err := strconv.Atoi(limitQuery)
	if err != nil || limit < 1 || limit > 100 {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Error: "invalid limit parameter: must be an integer between 1 and 100",
		})
	}

	bookIDStr := fiberCtx.Params("book_id")

	bookID, err := strconv.ParseInt(bookIDStr, 10, 64)
	if err != nil || bookID < 1 {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Error: "invalid book_id parameter",
		})
	}

	comments, total, err := t.service.GetCommentsByBookID(fiberCtx.Context(), int64(page), int64(limit), bookID)
	if err != nil {
		return fmt.Errorf("GetCommentsByBookId(%d, %d, %d): %w", page, limit, bookID, err)
	}

	response := ConvertToGetCommentsResponse(comments, total)

	return fiberCtx.Status(fiber.StatusOK).JSON(response)
}

func (t *Transport) GetCommentsByUserID(fiberCtx *fiber.Ctx) error {
	pageQuery := fiberCtx.Query("page")
	if pageQuery == "" {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Error: "page parameter is required",
		})
	}

	page, err := strconv.Atoi(pageQuery)
	if err != nil || page < 0 {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Error: "invalid page parameter: must be a non negative integer",
		})
	}

	limitQuery := fiberCtx.Query("limit")
	if limitQuery == "" {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Error: "limit parameter is required",
		})
	}

	limit, err := strconv.Atoi(limitQuery)
	if err != nil || limit < 1 || limit > 100 {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Error: "invalid limit parameter: must be an integer between 1 and 100",
		})
	}

	userIDStr := fiberCtx.Params("user_id")

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil || userID < 1 {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Error: "invalid user_id parameter",
		})
	}

	comments, total, err := t.service.GetCommentsByUserID(fiberCtx.Context(), int64(page), int64(limit), userID)
	if err != nil {
		return fmt.Errorf("GetCommentsByUserId(%d, %d, %d): %w", page, limit, userID, err)
	}

	response := ConvertToGetCommentsResponse(comments, total)

	return fiberCtx.Status(fiber.StatusOK).JSON(response)
}
