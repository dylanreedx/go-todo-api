package controllers

import (
	"net/http/httptest"
	"testing"

	"github.com/carbondesigned/go-todo/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetAllTodos(t *testing.T) {
	/* 	type args struct {
		c *fiber.Ctx
	} */
	tests := []struct {
		description  string
		route        string
		expectedCode int
		method       string
	}{
		{
			description:  "successfully fetch all todos",
			route:        "/",
			expectedCode: 200,
			method:       "GET",
		},
		{
			description:  "successfully update a todo",
			route:        "/:id",
			expectedCode: 201,
			method:       "PUT",
		},
	}
	app := fiber.New()

	routes.SetupRoutes(app)
	// Iterate through test single test cases
	for _, test := range tests {
		// Create a new http request with the route from the test case
		req := httptest.NewRequest(test.method, test.route, nil)

		// Perform the request plain with the app,
		// the second argument is a request latency
		// (set to -1 for no latency)
		resp, _ := app.Test(req, 1)

		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
