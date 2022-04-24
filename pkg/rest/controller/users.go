package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/radekkrejcirik01/radek-golang-backend/pkg/database"
	"github.com/radekkrejcirik01/radek-golang-backend/pkg/model/users"
)

// UsersGet GET /users
func UsersGet(c *fiber.Ctx) error {
	id := c.Params("id")
	t := &users.USERS{}
	if err := users.Read(database.DB, t, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(resp{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(resp{
		Status: "ok",
		Data:   &[]users.USERS{*t},
	})
}

func UsersGetAll(c *fiber.Ctx) error {
	t := &[]users.USERS{}
	if err := users.ReadAll(database.DB, t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(resp{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(resp{
		Status: "ok",
		Data:   t,
	})
}

// UsersPost POST /users
func UsersPost(c *fiber.Ctx) error {
	t := &users.USERS{}
	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(resp{
			Status:  "error",
			Message: err.Error(),
		})
	}
	if err := users.Create(database.DB, t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(resp{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(resp{Status: "ok"})
}

// UsersPut PUT /users
func UsersPut(c *fiber.Ctx) error {
	t := &users.USERS{}
	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(resp{
			Status:  "error",
			Message: err.Error(),
		})
	}
	if err := users.Update(database.DB, t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(resp{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(resp{Status: "ok"})
}

// UsersDel DELETE /users/:id
func UsersDel(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := users.DeleteByID(database.DB, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(resp{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(resp{Status: "ok"})
}
