package routes_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Kevindm14/docucenter-test/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// mock database
type MockDB struct {
	mock.Mock
}

func TestRoutes(t *testing.T) {
	// Structure for specific routes tests
	tests := []struct {
		name         string
		route        string
		expectedCode int
	}{
		{
			name:         "Get customers",
			route:        "/api/v1/customers",
			expectedCode: 200,
		},
		{
			name:         "Get ground deliveries",
			route:        "/api/v1/ground-deliveries",
			expectedCode: 200,
		},
		{
			name:         "Get maritime deliveries",
			route:        "/api/v1/maritime-deliveries",
			expectedCode: 200,
		},
	}

	// Run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Create request
			req := httptest.NewRequest(http.MethodGet, test.route, nil)

			// Create new fiber app
			app := fiber.New()

			// Set routes
			routes.SetRoutesApiV1(app, nil)

			// Process request
			resp, _ := app.Test(req, 1)

			assert.Equalf(t, test.expectedCode, resp.StatusCode, test.name)
		})
	}
}
