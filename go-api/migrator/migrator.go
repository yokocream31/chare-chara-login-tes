package main

import(
	"context"
	"fmt"

	"back-challe-chara2022/db"
	"back-challe-chara2022/entity/db_entity"

	"go.mongodb.org/mongo-driver/bson/primitive"

)

// DBマイグレート
func main() {

	// DBの初期化
	db.InitDB()

	// test success
	bearCollection := db.MongoClient.Database("insertDB").Collection("bears")
	doc := &db_entity.Bear{
		BearId: primitive.NewObjectID(),
		BearName: "kichiro",
        BearIcon: "img_dir/test.png",
        Detail: "テスト",
	}
	_, err := bearCollection.InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Println("Error inserting bear")
        panic(err)
    } else {
		fmt.Println("Successfully inserting bear")
	}
}