package repository

import (
	"Chatting/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

var cntx context.Context = context.TODO()

type UserRepository interface {
	PostUser(client *mongo.Client,user *model.User) (interface{}, error)
	GetOneUser(client *mongo.Client,filter bson.M) (*model.User, error)
}



func GetOneUser(client *mongo.Client,filter bson.M)(*model.User, error){
	ctx, cancel := context.WithTimeout(context.Background(),time.Minute *5)
	defer cancel()
	var user model.User
	collection := client.Database("webboard").Collection("users")
	userReturned := collection.FindOne(ctx,filter)
	if err := userReturned.Decode(&user); err != nil{
		log.Println("Error retrieving user")
		return &model.User{},err
	}
	log.Println(user.Name, user.Password)
	return &user,nil
}

func PostUser(client *mongo.Client,user *model.User) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(),time.Minute * 5)
	defer cancel()
	collection := client.Database("webboard").Collection("users")
	insertResult, err := collection.InsertOne(ctx, user)
	if err != nil{
		log.Println("Error on Inserting new user", err)
		return "", err
	}
	return insertResult.InsertedID, nil
}