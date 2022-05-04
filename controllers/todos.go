package controllers

import (
	"context"
	"fmt"

	"github.com/carbondesigned/go-todo/db"
	"github.com/carbondesigned/go-todo/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

var todoCollection = db.MongoClient().Database("todo-app").Collection("todos")

func AddTodo(c *fiber.Ctx) error {
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

func GetAllTodos(c *fiber.Ctx) error {
	cursor, err := todoCollection.Find(context.TODO(), bson.D{})

	if err != nil {
		panic(err)
	}
	var todos []bson.M

	if err = cursor.All(context.TODO(), &todos); err != nil {
		panic(err)
	}

	return c.JSON(todos)
}
