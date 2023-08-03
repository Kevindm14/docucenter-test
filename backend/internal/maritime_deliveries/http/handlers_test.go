package http_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
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

func TestMaritimeDeliveriesEndPoints(t *testing.T) {
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
	}

	db.Create(&customers)

	id := customers[0].ID

	maritimeDeliveries := []models.MaritimeDelivery{
		{
			CustomerID:      int(id),
			ProductQuantity: 1,
			ProductType:     "type 1",
			DeliveryPort:    "port 1",
			VehiclePlate:    "",
			GuideNumber:     "guide 1",
			ShippingPrice:   1000,
		},
		{
			CustomerID:      int(id),
			ProductQuantity: 1,
			ProductType:     "type 2",
			DeliveryPort:    "port 2",
			VehiclePlate:    "",
			GuideNumber:     "guide 1",
			ShippingPrice:   1000,
		},
	}

	db.Create(&maritimeDeliveries)

	// Structure for specific routes tests
	tests := []struct {
		name         string
		method       string
		route        string
		body         string
		expectedCode int
	}{
		{
			method:       http.MethodGet,
			name:         "Get Maritime Deliveries",
			route:        "/api/v1/maritime-deliveries",
			body:         "",
			expectedCode: fiber.StatusOK,
		},
		{
			method: http.MethodPost,
			name:   "Create Maritime Delivery",
			route:  "/api/v1/maritime-deliveries",
			body: `{
				"customer_id": ` + strconv.FormatUint(uint64(id), 10) + `,
				"product_quantity": 1,
				"product_type": "type 3",
				"port": "port 3",
				"vehicle_plate": "ABC1234Z",
				"guide_number": "guide 3",
				"shipping_price": 1000
				"registration_date": "2021-09-01T00:00:00Z"
				"delivery_date": "2021-09-01T00:00:00Z"
			}`,
			expectedCode: fiber.StatusCreated,
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

			// Delete all customers and maritime deliveries
			db.Delete(&customers)
			db.Delete(&maritimeDeliveries)
		})
	}
}
