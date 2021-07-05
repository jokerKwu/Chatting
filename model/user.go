package model

type User struct {
	Name  string `json:"name" xml:"name" bson:"name" validate:"required"`
	Password string `json:"password" bson:"password" validate:"required"`
}