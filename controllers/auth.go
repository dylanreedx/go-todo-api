package controllers

import (
	"time"

	"github.com/carbondesigned/go-todo/models"
	"github.com/carbondesigned/go-todo/utils"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func AuthRequired() func(c *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: utils.Secret(),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	})
}

func Signup(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}
	password := user.Password
	email := user.Email
	username := user.Username

	ctx, cancel := utils.Context()
	defer cancel()
	if err := c.BodyParser(user); err != nil {
		return err
	}

	// check if user already exists
	var userFound models.User
	err := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&userFound)
	if err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"success": false,
			"message": "User already exists",
		})
	}

	err = userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&userFound)
	if err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"success": false,
			"message": "User already exists",
		})
	}

	// simple password length validation
	if len(password) < 6 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Password must be at least 6 characters",
		})
	}

	// validation of username
	if len(username) < 4 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Username must be at least 4 characters",
		})
	}

	// hash password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	result, err := userCollection.InsertOne(ctx, user)

	if err != nil {
		panic(err)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "User Created",
		"data":    result,
	})
}

// It takes the email and password from the request body, checks if the user exists in the database, if
// it does, it compares the password from the request body with the hashed password from the database,
// if they match, it creates a JWT token and sends it back to the client
func Signin(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	email := user.Email
	password := user.Password

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

	// compare hashed password from db with password from post body
	if !utils.CheckPasswordHash(password, userFound.Password) || userFound.Email != email {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Wrong Credentials",
		})
	}
	// Create the Claims
	claims := jwt.MapClaims{
		"id":    userFound.ID,
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
