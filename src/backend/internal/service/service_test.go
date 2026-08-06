package service

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type stubHealthService struct{}

func (s stubHealthService) Healthz() string {
	return "ok"
}

type stubFooBarService struct{}

func (s stubFooBarService) FooBar() string {
	return "foo bar"
}

func TestDefaultServiceHealthz(t *testing.T) {
	t.Parallel()

	appService := NewDefaultService(stubHealthService{}, stubFooBarService{})

	if got := appService.Healthz(); got != "ok" {
		t.Fatalf("Healthz() = %q, want %q", got, "ok")
	}
}

func TestDefaultServiceFooBar(t *testing.T) {
	t.Parallel()

	appService := NewDefaultService(stubHealthService{}, stubFooBarService{})

	if got := appService.FooBar(); got != "ok -> foo bar" {
		t.Fatalf("FooBar() = %q, want %q", got, "ok -> foo bar")
	}
}

func TestNewDefaultHealthService(t *testing.T) {
	t.Parallel()

	healthService := NewDefaultHealthService()
	require.NotNil(t, healthService)
	require.Equal(t, "ok", healthService.Healthz())
}

func TestNewDefaultFooBarService(t *testing.T) {
	t.Parallel()

	fooBarService := NewDefaultFooBarService()
	require.NotNil(t, fooBarService)
	require.Equal(t, "foo bar", fooBarService.FooBar())
}

func TestMockServiceMethods(t *testing.T) {
	t.Parallel()

	mockService := &MockService{}
	mockService.On("Healthz").Return("ok")
	mockService.On("FooBar").Return("foo bar")

	require.Equal(t, "ok", mockService.Healthz())
	require.Equal(t, "foo bar", mockService.FooBar())
	mockService.AssertExpectations(t)
}
