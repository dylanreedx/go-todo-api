package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title    string             `bson:"title,omitempty" json:"title,omitempty"`
	Complete bool               `bson:"complete" json:"complete"`
	UserId   string             `json:"userId,omitempty" bson:"userId,omitempty"`
}
