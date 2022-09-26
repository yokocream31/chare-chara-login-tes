package user_controller

import (
	"back-challe-chara2022/entity/request_entity/body"
	"back-challe-chara2022/db"
	// "back-challe-chara2022/entity/db_entity"

	"net/http"
	"fmt"
	"io/ioutil"
	"context"
	
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gin-gonic/gin"
)

type UserController struct {}

type UserStatusResponse struct {
	StatusId string `json:"statusId"`
}

type UserIconResponse struct {
	UserIcon []byte `json:"userIcon"`
}

// PATCH: /user/status/<uuid: user_id>
func (uc UserController) PatchUserStatus(c *gin.Context) {

	// スタンプが押された際に，userのステータスを更新

	user_id, _ := primitive.ObjectIDFromHex(c.Param("user_id"))
	fmt.Println(user_id) // debug message

	var request body.PatchUserStatusBody
	// bodyのjsonデータを構造体にBind
	if err := c.Bind(&request); err != nil {
		// bodyのjson形式が合っていない場合
		c.JSON(http.StatusBadRequest, gin.H{"result": err.Error()})
		return
	}

	// userCollection := db.MongoClient.Database("insertDB").Collection("users")
	// var doc db_entity.User
	// // 検索条件
	// filter := bson.D{{"UserId", user_id}}
	// // query the database
	// if err := userCollection.FindOne(context.TODO(), filter).Decode(&doc); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"result": err.Error()})
	// 	return
	// } else if err == mongo.ErrNoDocuments {
	// 	fmt.Printf("No document was found with the user_id")
	// 	c.JSON(http.StatusOK, gin.H{})
	// 	return
	// }

	response := UserStatusResponse{StatusId: "ぬまり中"}
	c.JSON(http.StatusOK, response)
	return
}

// GET: /user/community/<uuid: user_id>
// $inを用いることで1つのクエリでいけるかも
func (uc UserController) GetUserCommunity(c *gin.Context) {

	// user_idが所属するコミュニティのcommunity_nameを全て返す

	user_id, _ := primitive.ObjectIDFromHex(c.Param("user_id"))
	fmt.Println(user_id) // debug message

	var err error

	userCollection := db.MongoClient.Database("insertDB").Collection("users")
	var doc_filter bson.M
	// 検索条件
	filter := bson.D{{"_id", user_id}}
	// query the user collection
	err = userCollection.FindOne(context.TODO(), filter, options.FindOne().SetProjection(bson.M{"communityId": 1})).Decode(&doc_filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": err.Error()})
		return
	} else if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the user_id")
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	var doc_community bson.M
	// 検索条件
	filter_community := bson.M{"_id": doc_filter["communityId"]}
	// query the community collection
	communityCollection := db.MongoClient.Database("insertDB").Collection("communities")
	err = communityCollection.FindOne(context.TODO(), filter_community, options.FindOne().SetProjection(bson.M{"communityName": 1, "_id": 0})).Decode(&doc_community)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": err.Error()})
		return
	} else if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the user_id")
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	c.JSON(http.StatusOK, doc_community)
	return


}

// GET: /user/icon/<uuid: user_id>
func (uc UserController) GetUserIcon(c *gin.Context) {

	// user_idのユーザのiconを返す

	user_id, _ := primitive.ObjectIDFromHex(c.Param("user_id"))
	fmt.Println(user_id) // debug message

	// 画像のbyteデータ読み込み
	buf, err := ioutil.ReadFile("img_dir/test.png")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": err.Error()})
		return
	} else {
		response := UserIconResponse{UserIcon: buf}
		c.JSON(http.StatusOK, response)
	}

}