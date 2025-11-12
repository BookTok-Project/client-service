package transport

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Transport struct {
	requestReader requestReader
	service       service
}

func New(requestReader requestReader, service service) *Transport {
	return &Transport{
		requestReader: requestReader,
		service:       service,
	}
}

func (t *Transport) HelloWorld(fiberCtx *fiber.Ctx) error {
	return fiberCtx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Hello": "World!",
	})
}

func (t *Transport) GetSubscribeeIDs(fiberCtx *fiber.Ctx) error {
	subscriberIDString := fiberCtx.Params("subscriber_id")

	subscriberID, err := strconv.ParseInt(subscriberIDString, 10, 64)
	if err != nil {
		return fmt.Errorf("parse int: %w", err)
	}

	subscribeeIDs, err := t.service.GetSubscribeeIDs(fiberCtx.Context(), subscriberID)
	if err != nil {
		return fmt.Errorf("get subscribers: %w", err)
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(subscribeeIDs)
}

func (t *Transport) GetSubscriberIDs(fiberCtx *fiber.Ctx) error {
	subscribeeIDString := fiberCtx.Params("subscribee_id")

	subscribeeID, err := strconv.ParseInt(subscribeeIDString, 10, 64)
	if err != nil {
		return fmt.Errorf("parse int: %w", err)
	}

	subscriberIDs, err := t.service.GetSubscriberIDs(fiberCtx.Context(), subscribeeID)
	if err != nil {
		return fmt.Errorf("get subscribers: %w", err)
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(subscriberIDs)
}

func (t *Transport) Subscribe(fiberCtx *fiber.Ctx) error {
	var req subscribeRequest

	if err := t.requestReader.ReadAndValidateFiberBody(fiberCtx, &req); err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Error: err.Error(),
		})
	}

	err := t.service.Subscribe(fiberCtx.Context(), req.SubscriberID, req.SubscribeeID)
	if err != nil {
		return err
	}

	return fiberCtx.SendStatus(fiber.StatusCreated)
}
