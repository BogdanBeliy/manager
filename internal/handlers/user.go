package handlers

import "github.com/gofiber/fiber/v2"

func UserHandlers(router fiber.Router) {
	router.Post("/", Create)
}

func Create(ctx *fiber.Ctx) error {
	if err := ctx.SendString("CREATE"); err != nil {
		return err
	}
	return nil
}
