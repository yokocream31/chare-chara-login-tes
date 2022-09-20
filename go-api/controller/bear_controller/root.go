package bear_controller

import (
	"back-challe-chara2022/entity/request_entity/body"

	"net/http"
	"fmt"
	"time"
	"math/rand"
	"io/ioutil"
	
	"github.com/gin-gonic/gin"
)

type BearController struct {}

type BearResponse struct {
	Response string `json:"response"`
}

type BearHistoryResponse struct {
	Message []string `json:"message"`
}

type BearCustomResponse struct {
	BearIcon []byte `json:"bearIcon"`
	BearTone uint `json:"bearTone"`
}



// POST: /bear/<str: user_id>
func (bc BearController) PostResponse(c *gin.Context) {
	
	// 送られてきた内容（message）はDBに保存
	// ランダムにresponseを返却

	var request body.PostSendBearBody
	// bodyのjsonデータを構造体にBind
	if err := c.Bind(&request); err != nil {
		// bodyのjson形式が合っていない場合
		c.JSON(http.StatusBadRequest, gin.H{"result": err.Error()})
		return
	}

	user_id := c.Param("user_id")
	fmt.Println(user_id) // debug message

	// 仮レスポンス
	response := [8]string{"そうだよね", "もう一回最初から教えてよ", "そこってどういうことなの？", 
						"頑張ってるやん！", "もうちょっと詰めてみようよ！", "君がそんなに考えてわかんないなら，誰もわかんないよ！", 
						"今，頭が回らないだけで少し時間を空けて考えたらわかる時もあるよ！", "それはもう心が１回休めって言ってるんだよ"}

	rand.Seed(time.Now().UnixNano())
    var idx int = rand.Intn(8)
	fmt.Println(idx)

	talk := BearResponse{Response: response[idx]}

	c.JSON(http.StatusOK, talk)

}

// GET: /bear/history/<uuid:user_id>
func (bc BearController) GetHistory(c *gin.Context) {

	// 指定されたuser_idのユーザのクマとの対話履歴を返す

	user_id := c.Param("user_id")
	fmt.Println(user_id) // debug message

	message := []string{"aaaaaa", "iiiii", "uuuuu", "eeeee", "ooooo"}

	response := BearHistoryResponse{Message: message}

	c.JSON(http.StatusOK, response)
}

// GET: /bear/custom/<uuid: user_id>
func (bc BearController) GetCustom(c *gin.Context) {
	
	// user_idのユーザの，クマのiconデータ，口調idを返す

	user_id := c.Param("user_id")
	fmt.Println(user_id) // debug message

	// 画像のbyteデータ読み込み
	buf, err := ioutil.ReadFile("img_dir/test.png")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": err.Error()})
		return
	} else {
		response := BearCustomResponse{BearIcon: buf, BearTone: 1}
		c.JSON(http.StatusOK, response)
	}
}

