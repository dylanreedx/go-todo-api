package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/carbondesigned/go-todo/db"
	"github.com/carbondesigned/go-todo/models"
	"github.com/gofiber/fiber/v2"
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
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	filter := bson.M{}
	findOptions := options.Find()
	cursor, err := todoCollection.Find(ctx, filter, findOptions)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "todos Not found",
			"error":   err,
		})
	}
	for cursor.Next(ctx) {
		var catchphrase models.Todo
		cursor.Decode(&catchphrase)
		todos = append(todos, catchphrase)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": todos,
	})
}

func GetTodoById(c *fiber.Ctx) error {
	var todo models.Todo
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	objId, err := primitive.ObjectIDFromHex(c.Params("id"))
	findResult := todoCollection.FindOne(ctx, bson.M{"_id": objId})
	if err := findResult.Err(); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "todo not found",
			"error":   err,
		})
	}

	err = findResult.Decode(&todo)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "todo not found",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    todo,
		"success": true,
	})
}
