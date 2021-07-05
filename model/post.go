package model

import (
	"time"
)

type Post struct{
	ID      int       `json:"id" bson:"id" validate:"required"`
	Title   string    `json:"title" bson:"title" validate:"required"`
	Content string    `json:"content" bson:"content" validate:"required"`
	Author  string    `json:"author" bson:"author" validate:"required"`
	Date    time.Time `json:"date" bson:"date" validate:"omitempty"'`
}
