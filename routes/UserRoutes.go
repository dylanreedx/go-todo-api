package routes

import (
	"github.com/carbondesigned/go-todo/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(route fiber.Router) {
	route.Get("/", controllers.GetAllUsers)
	route.Get("/:id", controllers.GetUserById)
	route.Delete("/:id", controllers.DeleteUser)
}
