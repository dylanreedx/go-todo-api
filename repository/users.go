package repository

import (
	"errors"

	"github.com/carbondesigned/go-todo/db"
	"github.com/carbondesigned/go-todo/models"
	"github.com/carbondesigned/go-todo/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	GetById(c *fiber.Ctx) error
}

var UserCollection = db.MongoClient().Database("todo-app").Collection("users")

func GetUserById(c *fiber.Ctx) error {
	var user models.User
	ctx, cancel := utils.Context()
	defer cancel()
	objId, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return errors.New(err.Error())
	}
	findResult := UserCollection.FindOne(ctx, bson.M{"_id": objId})
	if err := findResult.Err(); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User not found",
			"error":   err,
		})
	}

	err = findResult.Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User not found",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    user,
		"success": true,
	})
}
