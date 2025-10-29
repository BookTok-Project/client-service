package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func registerRoutes(app *fiber.App, transport transport) {
	app.Use(recover.New())
	app.Get("/", transport.HelloWorld)
	app.Post("/subscribe", transport.Subscribe)
	app.Get("/subscribers/:subscribee_id", transport.GetSubscriberIDs)
	app.Get("/subscribees/:subscriber_id", transport.GetSubscribeeIDs)
	app.Post("/favorites", transport.AddFavorite)
	app.Delete("/favorites", transport.RemoveFavorite)
	app.Get("/favorites/:user_id", transport.ListFavorites)
}
