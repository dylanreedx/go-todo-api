package models

type Todo struct {
	Title    string `json:"title"`
	Complete bool   `json:"complete"`
}
