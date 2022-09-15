package user_controller

import (
	"back-challe-chara2022/entity/request_entity/body"

	"net/http"
	"fmt"
	"io/ioutil"
	
	"github.com/gin-gonic/gin"
)

type UserController struct {}

type UserStatusResponse struct {
	StatusId string `json:"statusId"`
}

type UserCommunityResponse struct {
	CommunityName []string `json:"communityName"`
}

type UserIconResponse struct {
	UserIcon []byte `json:"userIcon"`
}

// PATCH: /user/status/<uuid: user_id>
func (uc UserController) PatchUserStatus(c *gin.Context) {

	// スタンプが押された際に，userのステータスを更新

	user_id := c.Param("user_id")
	fmt.Println(user_id) // debug message

	var request body.PatchUserStatusBody
	// bodyのjsonデータを構造体にBind
	if err := c.Bind(&request); err != nil {
		// bodyのjson形式が合っていない場合
		c.JSON(http.StatusBadRequest, gin.H{"result": err.Error()})
		return
	}

	response := UserStatusResponse{StatusId: "ぬまり中"}
	c.JSON(http.StatusOK, response)

}

// GET: /user/community/<uuid: user_id>
func (uc UserController) GetUserCommunity(c *gin.Context) {

	// user_idが所属するコミュニティのcommunity_nameを全て返す

	user_id := c.Param("user_id")
	fmt.Println(user_id) // debug message

	// 仮定義
	community := []string{"Python", "Flutter", "荒川研", "高専ラボ", "Golang"}

	response := UserCommunityResponse{CommunityName: community}
	c.JSON(http.StatusOK, response)


}

// GET: /user/icon/<uuid: user_id>
func (uc UserController) GetUserIcon(c *gin.Context) {

	// user_idのユーザのiconを返す

	user_id := c.Param("user_id")
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