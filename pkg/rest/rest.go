package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/radekkrejcirik01/radek-golang-backend/pkg/rest/controller"
)

// Create new REST API serveer
func Create() *fiber.App {
	app := fiber.New()

	app.Get("/", controller.Index)
	app.Get("/users/:id", controller.UsersGet)
	app.Get("/users", controller.UsersGetAll)
	app.Post("/users", controller.UsersPost)
	app.Put("/users", controller.UsersPut)
	app.Delete("/users", controller.UsersDel)
	app.Get("/config", controller.Config)

	return app
}
