package routes

import (
	"github.com/carbondesigned/go-todo/controllers"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(route fiber.Router) {
	route.Post("/signup", controllers.Signup)
	route.Post("/signin", controllers.Signin)
	route.Post("/signout", controllers.Signout)
	route.Get("/", controllers.GetAllUsers)
	route.Get("/:id", controllers.GetUserById)
}
