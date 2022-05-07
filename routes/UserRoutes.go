package routes

import (
	"github.com/carbondesigned/go-todo/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(route fiber.Router) {
	route.Get("/", controllers.AuthRequired(), controllers.GetAllUsers)
	route.Get("/:id", controllers.AuthRequired(), controllers.GetUserById)
	route.Delete("/:id", controllers.AuthRequired(), controllers.DeleteUser)
}
