package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/carbondesigned/go-todo/db"
	"github.com/carbondesigned/go-todo/models"
	"github.com/carbondesigned/go-todo/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
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

// It takes the email and password from the request body, finds the user in the database, compares the
// password from the request body with the password from the database, and if they match, it creates a
// JWT token and sends it back to the client
func Signin(c *fiber.Ctx) error {
	pass := c.FormValue("password")
	email := c.FormValue("email")

	var userFound models.User
	ctx, cancel := utils.Context()
	defer cancel()

	err := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&userFound)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User not found",
			"error":   err,
		})
	}

	// compare password from post body with password from db
	if pass != userFound.Password || email != userFound.Email {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Credentials do not match",
			"error":   err,
		})
	}
	// Create the Claims
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString(utils.Secret())
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	c.Cookie(&fiber.Cookie{
		Name:  "token",
		Value: t,
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": email + " signed in",
		"token":   t,
	})
}

func Signout(c *fiber.Ctx) error {
	// sign out
	c.ClearCookie("token")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "User logged out",
	})
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
