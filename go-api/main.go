package main

import (
	. "back-challe-chara2022/SessionInfo"
	"back-challe-chara2022/routes"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var LoginInfo SessionInfo

func main() {
	router := gin.Default()

	//テンプレートの設定
	router.LoadHTMLGlob("**/view/*.html")

	// セッションの設定
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.GET("/login", routes.GetLogin)
	router.POST("/login", routes.PostLogin)

	menu := router.Group("/menu")
	menu.Use(sessionCheck())
	{
		menu.GET("/top", routes.GetMenu)
	}
	router.POST("/logout", routes.PostLogout)
	router.Run(":8080")

}

func sessionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {

		session := sessions.Default(c)
		LoginInfo.UserId = session.Get("UserId")

		// セッションがない場合、ログインフォームをだす
		if LoginInfo.UserId == nil {
			log.Println("ログインしていません")
			c.Redirect(http.StatusMovedPermanently, "/login")
			c.Abort() // これがないと続けて処理されてしまう
		} else {
			c.Set("UserId", LoginInfo.UserId) // ユーザidをセット
			c.Next()
		}
		log.Println("ログインチェック終わり")
	}
}
