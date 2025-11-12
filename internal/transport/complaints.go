package transport

import (
	"github.com/gofiber/fiber/v2"
)

func (t *Transport) GetComplaints(fiberCtx *fiber.Ctx) error {
	complaints, err := t.service.GetComplaints(fiberCtx.Context())
	if err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Error: "Something gone wrong taking complaints",
		})
	}
	response := ConvertToGetComplaintsResponse(complaints)
	return fiberCtx.Status(fiber.StatusOK).JSON(response)
}

func (t *Transport) AddComplaint(fiberCtx *fiber.Ctx) error {
	/*
	userIDStr := fiberCtx.Query("user_id")
	if userIDStr == "" {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Error: "user_id parameter is required",
		})
	}

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Error: "user_id must be integer",
		})
	}
	bookIDStr := fiberCtx.Query("book_id")
	if  bookIDStr == "" {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Error: "book_id parameter is required",
		})
	}
	bookID, err := strconv.ParseInt(bookIDStr, 10, 64)
	if err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Error: "book_id must be integer",
		})
	}

	text := fiberCtx.Query("text")
	if  text == "" {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Error: "text parameter is required",
		})
	}
	*/


	var req addComplaintRequest
	if err := t.requestReader.ReadAndValidateFiberBody(fiberCtx, &req); err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := t.service.AddComplaint(fiberCtx.Context(), req.UserID, req.BookID, req.Text); err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Error: "Adding complaint failed",
		})
	}

	// _ := ConvertToAddComplaintsResponse(complaint)
	return fiberCtx.Status(fiber.StatusCreated).JSON(fiber.Map{})
}
