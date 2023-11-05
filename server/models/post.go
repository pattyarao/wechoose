package models

type Post struct {
	Id      string   `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string   `json:"name" bson:"name"`
	Title   string   `json:"title" bson:"title"`
	Choices []Choice `json:"choices" bson:"choices"`
}

type Choice struct {
	Title string `json:"title" bson:"title"`
	Votes int    `json:"votes" bson:"votes"`
}
