package db

import (
	"context"
	"os"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

var MongoClient *mongo.Client


// Databaseの初期化
func InitDB() {

	// 空のコンテキストを作成
	ctx := context.Background()
	fmt.Println(os.Getenv("MONGO_URI")) // debug msg

	var err error
	// Create a new client and connect to the server
	MongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		panic(err)
	}
	// defer MongoClient.Disconnect(ctx)

	// Ping the server to DB
    err = MongoClient.Ping(ctx, readpref.Primary())
    if err != nil {
        fmt.Println("connection error:", err)
    } else {
        fmt.Println("connection success:")
    }

	// test : success
	// coll := MongoClient.Database("insertDB").Collection("haikus")
	// doc := bson.D{{"title", "Record of a Shriveled Datum"}, {"text", "No bytes, no problem. Just insert a document, in MongoDB"}}
	// _, err = coll.InsertOne(context.TODO(), doc)
	// if err != nil {
	// 	panic(err)
	// }
}