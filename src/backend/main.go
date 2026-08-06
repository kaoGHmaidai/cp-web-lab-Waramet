package main

import (
	"log"

	"github.com/gofiber/fiber/v3"

	router "cp-web-template-backend/internal/endpoints"
	"cp-web-template-backend/internal/service"
)

func main() {
	app := newApp(buildAppService())

	log.Fatal(app.Listen(":3000"))
}

func buildAppService() service.Service {
	healthService := service.NewDefaultHealthService()
	fooBarService := service.NewDefaultFooBarService()

	return service.NewDefaultService(healthService, fooBarService)
}

func newApp(appService service.Service) *fiber.App {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("ok")
	})

	router.Register(app,
		router.NewHealthzRouter(appService),
		router.NewFooBarRouter(appService),
	)

	return app
}
