package db

import (
	"context"
	"os"

	"github.com/carbondesigned/go-todo/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoClient() (client *mongo.Client) {
	// Any time you make requests to a server (the database, in this case), you should create a context using context.TODO() that the server will accept.
	// TODO fix .env variables
	utils.GetEnv()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_CONNECTION_URI")))
	if err != nil {
		panic(err)
	}
	return client
}
