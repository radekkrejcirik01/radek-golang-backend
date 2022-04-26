package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/radekkrejcirik01/radek-golang-backend/pkg/rest/controller"
)

// Create new REST API serveer
func Create() *fiber.App {
	app := fiber.New()

	app.Get("/", controller.Index)
	app.Get("/users/:id", controller.UserGet)
	app.Get("/users", controller.UsersGetAll)
	app.Put("/users", controller.UserPut)
	app.Delete("/users/:id", controller.UserDel)

	app.Post("/users/register", controller.UserPost)
	app.Post("/users/login", controller.UserLogin)

	app.Get("/config", controller.Config)

	return app
}
