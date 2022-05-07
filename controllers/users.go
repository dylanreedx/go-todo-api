package controllers

import (
	"github.com/carbondesigned/go-todo/db"
	"github.com/carbondesigned/go-todo/models"
	"github.com/carbondesigned/go-todo/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func GetAllUsers(c *fiber.Ctx) error {
	var users []models.User
	ctx, cancel := utils.Context()
	defer cancel()

	filter := bson.M{}
	findOptions := options.Find()

	cursor, err := userCollection.Find(ctx, filter, findOptions)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "cannot retrieve users",
			"error":   err,
		})
	}

	for cursor.Next(ctx) {
		var user models.User
		cursor.Decode(&user)
		users = append(users, user)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    users,
	})
}
func GetUserById(c *fiber.Ctx) error {
	var user models.User

	ctx, cancel := utils.Context()
	defer cancel()

	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Something went wrong with the id",
		})
	}
	err = userCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    user,
	})
}
func DeleteUser(c *fiber.Ctx) error {
	ctx, cancel := utils.Context()
	defer cancel()

	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Something went wrong with the id",
		})
	}
	_, err = userCollection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "User Deleted",
	})
}
