package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	cnst "github.com/kemper0530/go-handson/common"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

// FireBaseの設定ファイル読込
func SetUpFirebase() (auth *auth.Client, err error) {
	// FirebaseのSDKを使用するためのkeyを読み込み
	opt := option.WithCredentialsFile(GetFireBasePath())
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}
	auth, errAuth := app.Auth(context.Background())

	return auth, errAuth
}

// JWT検証
func AuthFirebase(c *gin.Context, auth *auth.Client) (result int, errMsg string) {

	// クライアントから送られてきた JWT 取得
	authHeader := c.GetHeader("Authorization")
	idToken := strings.Replace(authHeader, "Bearer ", "", 1)

	// JWT の検証
	token, err := auth.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		u := fmt.Sprintf("error verifying ID token: %v\n", err)
		log.Printf("error verifying ID token: %v\n", err)
		return cnst.JsonStatusNG, u
	}
	u := fmt.Sprintf("Verified ID token: %v\n", token)

	return cnst.JsonStatusOK, u
}

// firebase json path
func GetFireBasePath() string {
	// 環境変数の読込
	firebasejsonpath := os.Getenv("FIREBASE_PATH")
	return firebasejsonpath
}
