package healthz

import "github.com/gofiber/fiber/v3"

func (r Router) Register(app *fiber.App) {
	app.Get("/healthz", r.Healthz)
}
