package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/require"

	"cp-web-template-backend/internal/service"
)

func TestNewAppBaseEndpoint(t *testing.T) {
	t.Parallel()

	app := newApp(&service.MockService{})
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response, err := app.Test(request, fiber.TestConfig{Timeout: 0, FailOnTimeout: false})
	require.NoError(t, err)

	body, err := io.ReadAll(response.Body)
	require.NoError(t, err)
	require.NoError(t, response.Body.Close())

	require.Equal(t, http.StatusOK, response.StatusCode)
	require.Equal(t, "ok", string(body))
}

func TestBuildAppService(t *testing.T) {
	t.Parallel()

	appService := buildAppService()
	require.NotNil(t, appService)
	require.Equal(t, "ok", appService.Healthz())
	require.Equal(t, "ok -> foo bar", appService.FooBar())
}
