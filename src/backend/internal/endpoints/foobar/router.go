package foobar

import "github.com/gofiber/fiber/v3"

func (r Router) Register(app *fiber.App) {
	app.Get("/foo/bar", r.FooBar)
}
