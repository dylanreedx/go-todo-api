package db

import (
	"context"
	"os"

	"github.com/carbondesigned/go-todo/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// It connects to the MongoDB database using the connection URI stored in the environment variable MONGODB_CONNECTION_URI
func MongoClient() (client *mongo.Client) {
	utils.LoadEnv()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_CONNECTION_URI")))
	if err != nil {
		panic(err)
	}
	return client
}
