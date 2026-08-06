package router

import (
	"github.com/gofiber/fiber/v3"

	"cp-web-template-backend/internal/endpoints/foobar"
	"cp-web-template-backend/internal/endpoints/greet"
	"cp-web-template-backend/internal/endpoints/healthz"
	"cp-web-template-backend/internal/service"
)

type Registrar interface {
	Register(app *fiber.App)
}

func Register(app *fiber.App, registrars ...Registrar) {
	for _, registrar := range registrars {
		registrar.Register(app)
	}
}

type HealthzRouter struct {
	Service service.Service
}

func NewHealthzRouter(service service.Service) HealthzRouter {
	return HealthzRouter{Service: service}
}

func (r HealthzRouter) Register(app *fiber.App) {
	healthz.NewRouter(r.Service).Register(app)
}

type FooBarRouter struct {
	Service service.Service
}

func NewFooBarRouter(service service.Service) FooBarRouter {
	return FooBarRouter{Service: service}
}

func (r FooBarRouter) Register(app *fiber.App) {
	foobar.NewRouter(r.Service).Register(app)
}

type GreetRouter struct {
	Service service.Service
}

func NewGreetRouter(service service.Service) GreetRouter {
	return GreetRouter{Service: service}
}

func (r GreetRouter) Register(app *fiber.App) {
	greet.NewRouter(r.Service).Register(app)
}
