package delivery

import "github.com/gofiber/fiber/v2"

type transport interface {
	HelloWorld(fiberCtx *fiber.Ctx) error
}
