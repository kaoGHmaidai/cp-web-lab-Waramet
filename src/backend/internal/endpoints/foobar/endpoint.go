package foobar

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

func (r Router) FooBar(c fiber.Ctx) error {
	return c.SendString(r.Service.FooBar())
}
