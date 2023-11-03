package models

type Post struct {
	Id      string   `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string   `json:"name" bson:"name"`
	Title   string   `json:"title" bson:"title"`
	Choices []string `json:"choices" bson:"choices"`
}
