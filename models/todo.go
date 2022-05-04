package models

type Todo struct {
	Title    string `bson:"title" json:"title"`
	Complete bool   `bson:"complete" json:"complete"`
}
