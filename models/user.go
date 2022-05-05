package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email    string             `bson:"email" json:"email"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"password"`
}
