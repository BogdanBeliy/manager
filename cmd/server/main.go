package main

import (
	"github.com/BogdanBeliy/manager/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New(
		logger.Config{
			Format:        "[${ip}]:${port} ${status} - ${method} ${path}\n",
			TimeFormat:    "02-Jan-2006",
			TimeZone:      "Europe/Minsk",
			DisableColors: false,
		}))

	usersGroup := app.Group("users/")

	handlers.UserHandlers(usersGroup)

	if err := app.Listen(":8081"); err != nil {
		panic(err)
	}
}
