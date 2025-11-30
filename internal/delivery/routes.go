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
	app.Post("/comments", transport.AddComment)
	app.Get("/comments/book/:book_id", transport.GetCommentsByBookID)
	app.Get("/comments/user/:user_id", transport.GetCommentsByUserID)
	app.Post("/complaint", transport.AddComplaint)
	app.Get("/complaint", transport.GetComplaints)
	app.Post("/likes", transport.AddLike)
	app.Delete("/likes", transport.RemoveLike)
	app.Get("/likes/:user_id", transport.ListLike)
}
