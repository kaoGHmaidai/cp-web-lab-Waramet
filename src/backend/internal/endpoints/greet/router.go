package greet

import "github.com/gofiber/fiber/v3"

func (r Router) Register(app *fiber.App) {
	app.Get("/greet", r.Greet)
}
