package repository

import (
	m "Chatting/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)


type PostRepository interface{
	GetAllPost(client *mongo.Client, filter bson.M)(Posts, error)
	SavePost(client *mongo.Client, post m.Post) (interface{}, error)
	GetOnePost(client *mongo.Client, filter bson.M)(m.Post, error)
	UpdatePost(client *mongo.Client, updateData interface{}, filter bson.M) (int64, error)
	DeletePost(client *mongo.Client, filter bson.M)(int64 error)
}

type Posts []m.Post

func GetAllPost(client *mongo.Client, filter bson.M) (Posts, error){
	ctx, cancel := context.WithTimeout(context.Background(),time.Minute * 5)
	defer cancel()
	collection := client.Database("webboard").Collection("posts")
	cur, err := collection.Find(ctx, filter)
	if err != nil{
		log.Println("Error on Finding all the documents", err)
		return Posts{}, err
	}
	var posts Posts
	for cur.Next(ctx) {
		var post m.Post
		err = cur.Decode(&post)
		if err != nil {
			log.Println("Error on Decoding the document", err)
			return Posts{}, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
func SavePost(client *mongo.Client, post m.Post)(interface{}, error){
	ctx, cancel := context.WithTimeout(context.Background(),time.Minute * 5)
	defer cancel()
	log.Println("여기 들어옴?")
	collection := client.Database("webboard").Collection("posts")
	insertResult, err := collection.InsertOne(ctx, post)
	if err != nil{
		log.Println("Error on inserting new post", err)
		return "", err
	}
	return insertResult.InsertedID, nil
}
func GetOnePost(client *mongo.Client, filter bson.M)(m.Post, error){
	ctx, cancel := context.WithTimeout(context.Background(),time.Minute * 5)
	defer cancel()
	var post m.Post
	collection := client.Database("webboard").Collection("posts")
	postReturned := collection.FindOne(ctx, filter)
	if err := postReturned.Decode(&post); err != nil{
		log.Println("Error retrieving post")
		return m.Post{}, err
	}
	return post, nil
}
func UpdatePost(client *mongo.Client, updateData interface{}, filter bson.M) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(),time.Minute * 5)
	defer cancel()
	collection := client.Database("webboard").Collection("posts")
	updateQuery:= bson.D{{Key:"$set",Value:updateData}}
	updateResult, err := collection.UpdateOne(ctx, filter, updateQuery)
	if err != nil {
		log.Println("Error on updating one post",err)
		return -1, err
	}
	return updateResult.ModifiedCount, nil
}
func DeletePost(client *mongo.Client, filter bson.M)(int64, error){
	ctx, cancel := context.WithTimeout(context.Background(),time.Minute * 5)
	defer cancel()
	collection := client.Database("webboard").Collection("posts")
	deleteResult, err := collection.DeleteOne(ctx, filter)
	if err != nil{
		log.Println("Error on deleting one post",err)
		return -1, err
	}
	return deleteResult.DeletedCount, nil
}





