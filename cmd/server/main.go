package main

import (
	"github.com/BogdanBeliy/manager/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	usersGroup := app.Group("users/")

	handlers.UserHandlers(usersGroup)

	if err := app.Listen(":8081"); err != nil {
		panic(err)
	}
}
