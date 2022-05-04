package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoClient() (client *mongo.Client) {
	// Any time you make requests to a server (the database, in this case), you should create a context using context.TODO() that the server will accept.
	// TODO fix .env variables
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://dylan:!Blackmonkeywhiterabbit@cluster0.mezvi.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}
	return client
}
