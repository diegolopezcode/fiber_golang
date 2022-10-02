package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/google/uuid"
)

type User struct {
	FirstName string
	LastName  string
	Id        string
}

func hanlderUser(c *fiber.Ctx) error {
	user := User{
		FirstName: "Diego",
		LastName:  "Lopez",
	}

	return c.Status(fiber.StatusOK).JSON(user)

}

func hanlderCreateUser(c *fiber.Ctx) error {
	user := User{}
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	user.Id = uuid.NewString()

	return c.Status(fiber.StatusOK).JSON(user)

}

func main() {
	app := fiber.New()

	//Middleware
	app.Use(logger.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Diego !")
	})

	app.Use(requestid.New())
	userGroup := app.Group("/user")

	userGroup.Get("", hanlderUser)
	userGroup.Post("", hanlderCreateUser)

	app.Listen(":8080")
}
