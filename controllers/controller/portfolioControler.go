package controller

import (
	"log"

	"github.com/gin-gonic/gin"

	// DBアクセス用モジュール
	db "github.com/kemper0530/go-handson/models/db"
)

// FetchAllMembers は メンバー情報を取得する
func FetchAllMembers(c *gin.Context) {
	resultProducts := db.FetchAllMembers()

	// URLへのアクセスに対してJSONを返す
	c.JSON(200, resultProducts)
}

// work情報を取得する
func FetchAllWorker(c *gin.Context) {
	resultProducts := db.FetchAllWorker()

	// URLへのアクセスに対してJSONを返す
	c.JSON(200, resultProducts)
}

// FetchLoginInfo は 指定したIDのパスワードを取得する
func FetchLoginInfo(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if len(username) == 0 || len(password) == 0 {
		log.Panic("Error nothing URL parameter!!")
	}

	resultProduct := db.FindLoginID(username, password)

	// URLへのアクセスに対してJSONを返す
	c.JSON(200, resultProduct)
}
