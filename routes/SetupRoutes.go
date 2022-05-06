package routes

import (
	"github.com/carbondesigned/go-todo/utils"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success":     true,
			"message":     "You are at the root endpoint 😉",
			"github_repo": "https://github.com/MikeFMeyer/catchphrase-go-mongodb-rest-api",
		})
	})
	api := app.Group("/api")
	TodoRoutes(api.Group("/todos"))
	AuthRoutes(api.Group("/auth"))

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: utils.Secret(),
	}))

}
