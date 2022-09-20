package server

import (
	"back-challe-chara2022/controller/bear_controller"
	"back-challe-chara2022/controller/user_controller"
	
	"net/http"
	"os"
	
	"github.com/gin-gonic/gin"
)

// 初期化
func Init() {

	// ルーティング
	r := setRouter()
	// Server Run (Port 8080)
	if err := r.Run(":" + os.Getenv("PORT")); err != nil {
		panic(err)
	}
}

// ルーティング設定
func setRouter() *gin.Engine {
	
	r := gin.Default()

	// ミドルウェアの設定
	r.Use(CORSMiddleware())

	//ルーティング
	bear_group := r.Group("bear")
	{
		ctrl := bear_controller.BearController{}
		// 熊の返答を返す
		bear_group.POST(":user_id", ctrl.PostResponse)
		// クマとの対話履歴を返す
		bear_group.GET("history/:user_id", ctrl.GetHistory)
		// クマのiconデータ，口調idを返す
		bear_group.GET("custom/:user_id", ctrl.GetCustom)
	}

	user_group := r.Group("user")
	{
		ctrl := user_controller.UserController{}
		// userのステータスを更新
		user_group.PATCH("status/:user_id", ctrl.PatchUserStatus)
		// userの所属するコミュニティを全て取得
		user_group.GET("community/:user_id", ctrl.GetUserCommunity)
		// userのアイコンを取得
		user_group.GET("icon/:user_id", ctrl.GetUserIcon)	
	}
	return r
}

// CORSリクエストのためのミドルウェア
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// リクエストの送信元の指定
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		// 資格情報（Cookie、認証ヘッダー、TLSクライアント証明書）の送信をOKするか
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		// リクエスト間に使用できるHTTPヘッダーを指定
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		// 使用できるメソッドを指定
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, DELETE, PATCH, OPTIONS")

		// OPTIONSメソッドは，指定されたURLまたはサーバーの許可されている通信オプションをリクエストする
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		// コンテキストに設定を書き込んだのでポインタを遷移（線形リスト）
		c.Next()
	}
}

