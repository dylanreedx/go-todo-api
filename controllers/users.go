package controllers

import (
	"context"
	"fmt"

	"github.com/carbondesigned/go-todo/db"
	"github.com/carbondesigned/go-todo/models"
	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	Signup(ctx *fiber.Ctx) error
	Signin(ctx *fiber.Ctx) error
	GetUser(ctx *fiber.Ctx) error
	GetAllUsers(ctx *fiber.Ctx) error
	/* PutUser(ctx *fiber.Ctx) error */
	DeleteUser(ctx *fiber.Ctx) error
}

var userCollection = db.MongoClient().Database("todo-app").Collection("users")

func Signup(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return err
	}

	result, err := userCollection.InsertOne(context.TODO(), user)

	if err != nil {
		panic(err)
	}
	fmt.Println(result.InsertedID)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "User Created",
		"data":    result,
	})
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
