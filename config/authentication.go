package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	cnst "github.com/kemper0530/go-handson/common"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

// FireBase認証確認
func AuthFirebase(c *gin.Context) (result int, errMsg string) {
	// FirebaseのSDKを使用するためのkeyを読み込み
	u := ""
	opt := option.WithCredentialsFile(GetFireBasePath())
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		u = fmt.Sprintf("error: %v", err)
		os.Exit(1)
		return cnst.JsonStatusNG, u
	}
	auth, err := app.Auth(context.Background())
	if err != nil {
		u = fmt.Sprintf("error: %v", err)
		return cnst.JsonStatusNG, u
	}

	// クライアントから送られてきた JWT 取得
	authHeader := c.GetHeader("Authorization")
	idToken := strings.Replace(authHeader, "Bearer ", "", 1)

	// JWT の検証
	token, err := auth.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		u = fmt.Sprintf("error verifying ID token: %v\n", err)
		log.Printf("error verifying ID token: %v\n", err)
		return cnst.JsonStatusNG, u
	}
	u = fmt.Sprintf("Verified ID token: %v\n", token)

	return cnst.JsonStatusOK, u
}

// firebase json path
func GetFireBasePath() string {
	// 環境変数ファイルの読込
	err := godotenv.Load(fmt.Sprintf("config/%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	firebasejsonpath := os.Getenv("FIREBASE_PATH")

	return firebasejsonpath
}
