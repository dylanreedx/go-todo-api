package controllers

import (
	"context"
	"errors"
	"fmt"

	"github.com/carbondesigned/go-todo/db"
	"github.com/carbondesigned/go-todo/models"
	"github.com/carbondesigned/go-todo/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// It's creating a variable called `todoCollection` that is equal to the `todos` collection in the `todo-app` database.
var todoCollection = db.MongoClient().Database("todo-app").Collection("todos")

// We create a new instance of the Todo model, parse the request body, insert the new Todo into the database, and return the request body
func CreateTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)

	if err := c.BodyParser(todo); err != nil {
		return err
	}

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)
	todo.UserId = id

	result, err := todoCollection.InsertOne(context.TODO(), todo)

	if err != nil {
		panic(err)
	}
	fmt.Println(result.InsertedID)
	return c.Send(c.Body())
}

// We're creating a context with a timeout of 10 seconds, then we're creating a filter and findOptions
// variable, then we're creating a cursor and decoding the data into the todos variable, and finally
// we're returning the data in a JSON format
func GetAllTodos(c *fiber.Ctx) error {
	var todos []models.Todo
	var todo models.Todo
	ctx, cancel := utils.Context()
	defer cancel()

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)
	todo.UserId = id

	// filter the todos with the userId
	filter := bson.M{"userId": id}
	// sort the todos by createdAt
	sort := options.Find().SetSort(bson.M{"createdAt": -1})
	// find the todos
	cursor, err := todoCollection.Find(ctx, filter, sort)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Todos Not found",
			"error":   err,
		})
	}
	for cursor.Next(ctx) {
		var todo models.Todo
		cursor.Decode(&todo)
		todos = append(todos, todo)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": todos,
	})
}

// We're using the `fiber.Ctx` struct to get the `id` parameter from the URL, then we're using the
// `primitive.ObjectIDFromHex` function to convert the `id` parameter to a `primitive.ObjectID` type.
// Then we're using the `FindOne` function to find the todo with the `id` parameter. If the todo is
// found, we're decoding the todo into the `todo` variable and returning it as a JSON response. If the todo is not found, we're returning a JSON response with an error message
func GetTodoById(c *fiber.Ctx) error {
	var todo models.Todo
	ctx, cancel := utils.Context()
	defer cancel()
	objId, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return errors.New(err.Error())
	}

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)
	todo.UserId = id

	filter := bson.M{"_id": objId, "userId": id}
	err = todoCollection.FindOne(ctx, filter).Decode(&todo)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Todo Not found",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    todo,
		"success": true,
	})
}

// It takes the id of the todo from the URL, converts it to an ObjectID, and then deletes the todo from the database
func DeleteTodo(c *fiber.Ctx) error {
	ctx, cancel := utils.Context()
	defer cancel()
	objId, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return errors.New(err.Error())
	}

	_, err = todoCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Todo failed to delete",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Todo deleted successfully",
	})
}

func UpdateTodo(c *fiber.Ctx) error {
	var todo = new(models.Todo)

	ctx, cancel := utils.Context()
	defer cancel()
	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}

	objId, err := primitive.ObjectIDFromHex(c.Params("id"))

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)
	todo.UserId = id

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Todo Not found",
			"error":   err,
		})
	}

	update := bson.M{
		"$set": todo,
	}

	_, err = todoCollection.UpdateOne(ctx, bson.M{"_id": objId, "userId": id}, update)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Todo failed to update",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Todo Updated",
	})

}
