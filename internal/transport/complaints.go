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
	var req addComplaintRequest
	if err := t.requestReader.ReadAndValidateFiberBody(fiberCtx, &req); err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	
	if err := t.service.AddComplaint(fiberCtx.Context(), req.UserID, req.BookID, req.Text); err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Error: "Adding co",
		})
	}

	// _ := ConvertToAddComplaintsResponse(complaint)
	return fiberCtx.Status(fiber.StatusCreated).JSON(fiber.Map{"response": "created"})
}
