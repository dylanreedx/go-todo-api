package models

type User struct {
	Id       string `bson:"_id" json:"id"`
	Email    string `bson:"email" json:"email"`
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
}
