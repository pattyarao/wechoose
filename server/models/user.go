package models

type User struct {
	Id       string   `json:"id,omitempty" bson:"_id,omitempty"`
	Username string   `json:"user_name" bson:"user_name"`
	Password string   `json:"password,omitempty" bson:"password"`
	Polls    []string `json:"polls" bson:"polls"`
}
