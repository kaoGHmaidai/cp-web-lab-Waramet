package router

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"cp-web-template-backend/internal/service"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/require"
)

func TestRegisterMountsAllRouters(t *testing.T) {
	t.Parallel()

	appService := &service.MockService{}
	appService.On("Healthz").Return("ok")
	appService.On("FooBar").Return("foo bar")

	app := fiber.New()
	Register(app,
		HealthzRouter{Service: appService},
		FooBarRouter{Service: appService},
	)

	healthResponse := performRequest(t, app, http.MethodGet, "/healthz")
	healthBody, err := io.ReadAll(healthResponse.Body)
	require.NoError(t, err)
	require.NoError(t, healthResponse.Body.Close())
	require.Equal(t, http.StatusOK, healthResponse.StatusCode)

	require.Equal(t, "ok", string(healthBody))

	fooResponse := performRequest(t, app, http.MethodGet, "/foo/bar")
	fooBody, err := io.ReadAll(fooResponse.Body)
	require.NoError(t, err)
	require.NoError(t, fooResponse.Body.Close())
	require.Equal(t, http.StatusOK, fooResponse.StatusCode)
	require.Equal(t, "foo bar", string(fooBody))

	appService.AssertExpectations(t)
}

func TestHealthzRouterRegister(t *testing.T) {
	t.Parallel()

	appService := &service.MockService{}
	appService.On("Healthz").Return("ok")

	app := fiber.New()
	HealthzRouter{Service: appService}.Register(app)

	response := performRequest(t, app, http.MethodGet, "/healthz")
	body, err := io.ReadAll(response.Body)
	require.NoError(t, err)
	require.NoError(t, response.Body.Close())

	require.Equal(t, http.StatusOK, response.StatusCode)
	require.Equal(t, "ok", string(body))
	appService.AssertExpectations(t)
}

func TestFooBarRouterRegister(t *testing.T) {
	t.Parallel()

	appService := &service.MockService{}
	appService.On("FooBar").Return("foo bar")

	app := fiber.New()
	FooBarRouter{Service: appService}.Register(app)

	response := performRequest(t, app, http.MethodGet, "/foo/bar")
	body, err := io.ReadAll(response.Body)
	require.NoError(t, err)
	require.NoError(t, response.Body.Close())

	require.Equal(t, http.StatusOK, response.StatusCode)
	require.Equal(t, "foo bar", string(body))
	appService.AssertExpectations(t)
}

func performRequest(t *testing.T, app *fiber.App, method, path string) *http.Response {
	t.Helper()

	req := httptest.NewRequest(method, path, nil)
	response, err := app.Test(req, fiber.TestConfig{Timeout: 0, FailOnTimeout: false})
	require.NoError(t, err)
	return response
}
