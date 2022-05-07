package routes

import (
	"github.com/carbondesigned/go-todo/controllers"
	"github.com/gofiber/fiber/v2"
)

func TodoRoutes(route fiber.Router) {
	route.Get("/", controllers.AuthRequired(), controllers.GetAllTodos)
	route.Post("/create", controllers.AuthRequired(), controllers.CreateTodo)
	route.Get("/:id", controllers.AuthRequired(), controllers.GetTodoById)
	route.Delete("/:id", controllers.AuthRequired(), controllers.DeleteTodo)
	route.Put("/:id", controllers.AuthRequired(), controllers.UpdateTodo)
}
