package routes_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Kevindm14/docucenter-test/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// mock database
type MockDB struct {
	mock.Mock
}

// Test routes
func TestRoutes(t *testing.T) {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

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
			name:         "Get customer by id",
			route:        "/api/v1/customers/1",
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
			// Create new fiber app
			app := fiber.New()

			// Set routes
			routes.SetRoutesApiV1(app)

			req := httptest.NewRequest(http.MethodGet, test.route, nil)
			resp, err := app.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, resp.StatusCode)
		})
	}
}
