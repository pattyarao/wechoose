package models

type User struct {
	Id       string `json:"id,omitempty" bson:"_id,omitempty"`
	Username string `json:"name"`
	Password string `json:"password"`
}
