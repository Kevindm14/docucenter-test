package http_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Kevindm14/docucenter-test/config"
	"github.com/Kevindm14/docucenter-test/internal/models"
	"github.com/Kevindm14/docucenter-test/libraries"
	"github.com/Kevindm14/docucenter-test/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestCustomersEndPoints(t *testing.T) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	// Database connection
	db := config.PgDBConnectionTest()

	customers := []models.Customer{
		{
			FirstName: "Kevin 1",
			LastName:  "Dominguez",
			Email:     "",
			Phone:     "1234567890",
		},
		{
			FirstName: "Kevin 2",
			LastName:  "Dominguez",
			Email:     "",
			Phone:     "1234567890",
		},
	}

	db.Create(&customers)

	// Structure for specific routes tests
	tests := []struct {
		name         string
		method       string
		route        string
		body         string
		expectedCode int
		contains     []string
	}{
		{
			method:       http.MethodGet,
			name:         "Get customers",
			route:        "/api/v1/customers",
			body:         "",
			expectedCode: fiber.StatusOK,
			contains:     []string{"Kevin 1", "Kevin 2"},
		},
		{
			method:       http.MethodGet,
			name:         "Get customer by id",
			route:        "/api/v1/customers/1",
			body:         "",
			expectedCode: fiber.StatusNotFound,
			contains:     []string{},
		},
		{
			method:       http.MethodPost,
			name:         "Create customer",
			route:        "/api/v1/customers",
			body:         `{"first_name": "Kevin 3", "last_name": "Dominguez", "email": "", "phone": "1234567890"}`,
			expectedCode: fiber.StatusCreated,
			contains:     []string{"Kevin 3"},
		},
	}

	// Run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Create new fiber app
			app := fiber.New()

			// Set routes
			routes.SetRoutesApiV1(app, db)

			req := httptest.NewRequest(test.method, test.route, nil)

			if test.method == http.MethodPost {
				req, _ = http.NewRequest(test.method, test.route, strings.NewReader(test.body))
				req.Header.Set("Content-Type", "application/json")
			}

			token, err := libraries.GenerateToken(1)
			assert.NoError(t, err)

			req.Header.Add("Authorization", "Bearer "+token)

			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer resp.Body.Close()

			assert.Equal(t, test.expectedCode, resp.StatusCode)

			// Delete all customers
			db.Delete(&customers)
		})
	}
}
