package routes

import (
	"github.com/carbondesigned/go-todo/controllers"
	"github.com/gofiber/fiber/v2"
)

func TodoRoutes(route fiber.Router) {
	route.Get("/", controllers.GetAllTodos)
	route.Post("/create", controllers.CreateTodo)
	route.Get("/:id", controllers.GetTodoById)
	route.Delete("/:id", controllers.DeleteTodo)
	route.Put("/:id", controllers.UpdateTodo)
}
