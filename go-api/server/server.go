package server

import (
	//"back-challe-chara2022/controller/sensor_controller"
	"net/http"
	"os"
	
	"github.com/gin-gonic/gin"
)

// 初期化
func Init() {

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
	// s := r.Group("sensor")
	// {
	// 	ctrl := sensor_controller.SensorController{}
	// 	// センサーの取得
	// 	s.GET(":sensor_id", ctrl.GetSensor)
	// 	// センサーの登録
	// 	s.POST("", ctrl.PostSensor)
	// 	// センサーデータの保存
	// 	s.POST("data", ctrl.PostSensorData)
	// 	// センサーデータの取得
	// 	s.GET("data", ctrl.GetSensorData)
	// }

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
