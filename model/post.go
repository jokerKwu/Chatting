package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Post struct{
	*PostInput `bson:",inline"`
	ID	primitive.ObjectID `josn:"id" xml:"id" bson:"_id,omitempty"`
}

type PostInput struct{
	Title string 	`json:"title" xml:"title" bson:"title" validate:"required"`
	Content string	`json:"content" xml:"content" bson:"content" validate:"required"`
	Author string	`json:"author" xml:"author" bson:"author" validate:"required"`
	Date time.Time	`json:"date" xml:"date" bson:"date" validate:"omitempty"`
}
