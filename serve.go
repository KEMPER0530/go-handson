package main

import (
	// ロギングを行うパッケージ
	"log"

	// Gin
	"github.com/gin-gonic/gin"

	// MySQL用ドライバ
	_ "github.com/jinzhu/gorm/dialects/mysql"

	// コントローラー
	controller "github.com/kemper0530/go-handson/controllers/controller"
)

func main() {
	// サーバーを起動する
	serve()
}

func serve() {
	// デフォルトのミドルウェアでginのルーターを作成
	// Logger と アプリケーションクラッシュをキャッチするRecoveryミドルウェア を保有しています
	router := gin.Default()

	// CORS対応
	router.Use(Cors())

	// ルーターの設定
	// ログインID、パスワードを返却する
	router.GET("/fetchLoginInfo", controller.FetchLoginInfo)

	// メンバー情報のJSONを返す
	router.GET("/fetchAllMembers", controller.FetchAllMembers)

	// work情報のJSONを返す
	router.GET("/fetchAllWorker", controller.FetchAllWorker)
	if err := router.Run(":8090"); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}

// Cross-Origin Resource Sharing (CORS) is a mechanism
// that uses additional HTTP headers to let a
// user agent gain permission to access selected resources from a server
// on a different origin /(domain) than the site currently in use.
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Next()
	}
}
