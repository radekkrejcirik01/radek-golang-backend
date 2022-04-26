package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/radekkrejcirik01/radek-golang-backend/pkg/database"
	"github.com/radekkrejcirik01/radek-golang-backend/pkg/model/users"
)

// UserGet GET /users/:id
func UserGet(c *fiber.Ctx) error {
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

// UserGetAll GET /users
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

// UserPost POST /users
func UserPost(c *fiber.Ctx) error {
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

// UserPut PUT /users
func UserPut(c *fiber.Ctx) error {
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

// UserDel DELETE /users/:id
func UserDel(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := users.DeleteById(database.DB, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(resp{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(resp{Status: "ok"})
}

// UserLogin AUTHENTICATE /login
func UserLogin(c *fiber.Ctx) error {
	t := &users.USERS{}
	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(resp{
			Status:  "error",
			Message: err.Error(),
		})
	}
	if err := users.Login(database.DB, t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(resp{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusInternalServerError).JSON(resp{
		Status:  "succes",
		Message: "User exists!",
	})
}
