package main

import (
	"context"
	"log"

	"github.com/carbondesigned/go-todo/db"
	"github.com/carbondesigned/go-todo/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	app := fiber.New()
	app.Use(cors.New())

	routes.SetupRoutes(app)

	// ping mongo
	if err := db.MongoClient().Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	app.Listen(":3001")
}
