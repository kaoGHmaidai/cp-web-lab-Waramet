package service

import "fmt"

type Service interface {
	Healthz() string
	FooBar() string
}

type HealthService interface {
	Healthz() string
}

type FooBarService interface {
	FooBar() string
}

type DefaultService struct {
	healthService HealthService
	fooBarService FooBarService
}

func NewDefaultService(healthService HealthService, fooBarService FooBarService) *DefaultService {
	return &DefaultService{
		healthService: healthService,
		fooBarService: fooBarService,
	}
}

func (s *DefaultService) Healthz() string {
	return s.healthService.Healthz()
}

func (s *DefaultService) FooBar() string {
	return fmt.Sprintf("%s -> %s", s.healthService.Healthz(), s.fooBarService.FooBar())
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
