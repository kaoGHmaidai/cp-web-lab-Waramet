package greet

import (
	"github.com/gofiber/fiber/v3"

	"cp-web-template-backend/internal/service"
)

type Router struct {
	Service service.Service
}

func NewRouter(service service.Service) Router {
	return Router{Service: service}
}

func (r Router) Greet(c fiber.Ctx) error {
	name := c.Query("name")

	return c.JSON(fiber.Map{
		"message": r.Service.Greet(name),
	})
}
