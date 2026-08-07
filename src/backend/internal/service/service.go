package service

import "fmt"

type Service interface {
	Healthz() string
	FooBar() string
	Greet(name string) string
}

type HealthService interface {
	Healthz() string
}

type FooBarService interface {
	FooBar() string
}

type GreetService interface {
	Greet(name string) string
}

type DefaultService struct {
	healthService HealthService
	fooBarService FooBarService
	greetService  GreetService
}

func NewDefaultService(healthService HealthService, fooBarService FooBarService, greetService GreetService) *DefaultService {
	return &DefaultService{
		healthService: healthService,
		fooBarService: fooBarService,
		greetService:  greetService,
	}
}

func (s *DefaultService) Healthz() string {
	return s.healthService.Healthz()
}

func (s *DefaultService) FooBar() string {
	return fmt.Sprintf("%s -> %s", s.healthService.Healthz(), s.fooBarService.FooBar())
}

func (s *DefaultService) Greet(name string) string {
	return s.greetService.Greet(name)
}

type DefaultHealthService struct{}

func NewDefaultHealthService() HealthService {
	return &DefaultHealthService{}
}

func (s *DefaultHealthService) Healthz() string {
	return "ok"
}

type DefaultFooBarService struct{}

func NewDefaultFooBarService() FooBarService {
	return &DefaultFooBarService{}
}

func (s *DefaultFooBarService) FooBar() string {
	return "foo bar"
}

const defaultGreetName = "friend"

type DefaultGreetService struct{}

func NewDefaultGreetService() GreetService {
	return &DefaultGreetService{}
}

func (s *DefaultGreetService) Greet(name string) string {
	if name == "" {
		name = defaultGreetName
	}

	return fmt.Sprintf("Hello, %s!", name)
}
