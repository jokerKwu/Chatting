package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func MongoConnection()(*mongo.Client, error){
	ctx, cancel := context.WithTimeout(context.Background(),time.Second * time.Duration(ConnectTime.(int64)))
	defer cancel()
	//Set clinent options
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoUrl.(string)))
	if err != nil{
		log.Fatal(err)
	}
	//Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil{
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB !!")
	return client, err
}
