package healthz

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/require"

	"cp-web-template-backend/internal/service"
)

func TestNewRouter(t *testing.T) {
	t.Parallel()

	router := NewRouter(&service.MockService{})
	require.NotNil(t, router)
}

func TestRouterHealthz(t *testing.T) {
	t.Parallel()

	appService := &service.MockService{}
	appService.On("Healthz").Return("ok")

	app := fiber.New()
	NewRouter(appService).Register(app)

	request := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	response, err := app.Test(request, fiber.TestConfig{Timeout: 0, FailOnTimeout: false})
	require.NoError(t, err)

	body, err := io.ReadAll(response.Body)
	require.NoError(t, err)
	require.NoError(t, response.Body.Close())

	require.Equal(t, http.StatusOK, response.StatusCode)
	require.Equal(t, "ok", string(body))
	appService.AssertExpectations(t)
}
