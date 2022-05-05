package routes

import (
	"github.com/carbondesigned/go-todo/controllers"
	"github.com/gofiber/fiber/v2"
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
}

func TodoRoutes(route fiber.Router) {
	route.Get("/", controllers.GetAllTodos)
	route.Post("/create", controllers.CreateTodo)
	route.Get("/:id", controllers.GetTodoById)
	route.Delete("/:id", controllers.DeleteTodo)
	route.Put("/:id", controllers.UpdateTodo)
}
