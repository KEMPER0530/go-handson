package main

import (
	"fmt"
	"log"
	"os"
	"time"

	// Gin
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// config
	config "github.com/kemper0530/go-handson/config"
	// common
	common "github.com/kemper0530/go-handson/common"
	// MySQL用ドライバ
	_ "github.com/jinzhu/gorm/dialects/mysql"

	// コントローラー
	controller "github.com/kemper0530/go-handson/controllers/controller"
)

func main() {
	// 環境変数ファイルの読込
	err := godotenv.Load(fmt.Sprintf("config/%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// firebaseSDKの読込
	auth, err := config.SetUpFirebase()
	if err != nil {
		log.Fatal("Error loading firebase-auth file")
	}
	// commonに格納する
	common.Auth = auth

	// ポートの取得
	PORT := os.Getenv("PORT")

	// サーバーを起動する
	router := serve()
	if err := router.Run(":" + PORT); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}

func serve() *gin.Engine {

	// デフォルトのミドルウェアでginのルーターを作成
	// Logger と アプリケーションクラッシュをキャッチするRecoveryミドルウェア を保有しています
	router := gin.Default()

	// 本番設定の場合
	if os.Getenv("GO_ENV") == "production" {
		// 環境変数を設定します.
		os.Setenv("GIN_MODE", "release")
		gin.SetMode(gin.ReleaseMode)
		router = gin.New()
	}
	// CORS設定
	router.Use(setCors())

	// ルーターの設定
	// ログインID、パスワードを返却する
	router.POST("/api/fetchLoginInfo", controller.FetchLoginInfo)

	// メンバー情報のJSONを返す
	router.GET("/api/fetchAllMembers", controller.FetchAllMembers)

	// work情報のJSONを返す
	router.GET("/api/fetchAllWorker", controller.FetchAllWorker)

	// クレジットカード情報を登録し、結果のJSONを返す
	router.POST("/api/fetchCreditInfoRegist", controller.FetchCreditInfoRegist)

	// お問合せフォーム内容を登録し、メールを送信するかつ結果のJSONを返す
	router.POST("/api/fetchSendMailRegist", controller.FetchSendMailRegist)

	// Goアプリのステータスを返却する
	router.GET("/api/actuaterHealth", controller.ActuaterHealth)

	// profile情報のJSONを返す
	router.GET("/api/fetchProfileInfo", controller.FetchProfileInfo)

	// アカウント情報を仮登録し、結果をJSONを返す
	router.POST("/api/fetchRegistAccount", controller.FetchRegistAccount)

	// 仮登録後にメール送信する結果をJSONを返す
	router.POST("/api/fetchRegistAccountMail", controller.FetchRegistAccountMail)

	// ログインIDを受取り、氏名とメールアドレスを返却する
	router.POST("/api/fetchMailAdrInfo", controller.FetchMailAdrInfo)

	// 仮パスワードのリンクを押下された場合の挙動
	//router.Static("/static/css", "./static/css")
	router.LoadHTMLGlob("templates/*.tmpl")
	router.GET("/api/fetchSignUpAccountMail", controller.FetchSignUpAccountMail)

	// NEWSAPIの記事を取得し、フロントへ返却する
	router.POST("/api/fetchNewsInfo", controller.FetchNewsInfo)

	// アクセスログを登録する
	router.POST("/api/fetchRegistAccessLog", controller.FetchRegistAccessLog)

	// Lambdaからリクエストされた内容を登録する
	//router.POST("/api/fetchRegistBounce", controller.FetchRegistBounce)

	return router
}

// Cross-Origin Resource Sharing (CORS) is a mechanism
// that uses additional HTTP headers to let a
// user agent gain permission to access selected resources from a server
// on a different origin /(domain) than the site currently in use.
// CORS for All origins, allowing:
// - PUT and PATCH methods
// - Origin header
// - Credentials share
// - Preflight requests cached for 1 hours
func setCors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "Authorization", "Accept", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Cache-Control", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           1 * time.Hour,
	})
}
