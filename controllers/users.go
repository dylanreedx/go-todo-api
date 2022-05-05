package controllers

import "github.com/gofiber/fiber/v2"

type AuthController interface {
	SignUp(ctx *fiber.Ctx) error
	SignIn(ctx *fiber.Ctx) error
	GetUser(ctx *fiber.Ctx) error
	GetAllUsers(ctx *fiber.Ctx) error
	/* PutUser(ctx *fiber.Ctx) error */
	DeleteUser(ctx *fiber.Ctx) error
}

func NewAuthController()

func Signup(c *fiber.Ctx) error {
	panic("not implemented")
}
func Signin(c *fiber.Ctx) error {
	panic("not implemented")
}
func Signout(c *fiber.Ctx) error {
	panic("not implemented")
}
func GetAllUsers(c *fiber.Ctx) error {
	panic("not implemented")
}
func GetUserById(c *fiber.Ctx) error {
	panic("not implemented")
}
func DeleteUser(c *fiber.Ctx) error {
	panic("not implemented")
}
