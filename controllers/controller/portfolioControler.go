package controller

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	// constクラス
	cnst "github.com/kemper0530/go-handson/common"
	// DBアクセス用モジュール
	db "github.com/kemper0530/go-handson/models/db"
)

// FetchAllMembers は メンバー情報を取得する
func FetchAllMembers(c *gin.Context) {
	resultProducts := db.FetchAllMembers()

	// URLへのアクセスに対してJSONを返す
	c.JSON(cnst.HttpStatusOK, resultProducts)
}

// work情報を取得する
func FetchAllWorker(c *gin.Context) {
	resultProducts := db.FetchAllWorker()

	// URLへのアクセスに対してJSONを返す
	c.JSON(cnst.HttpStatusOK, resultProducts)
}

// FetchLoginInfo は 指定したIDのパスワードを取得する
func FetchLoginInfo(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if len(username) == cnst.ZERO || len(password) == cnst.ZERO {
		log.Panic("Error nothing URL parameter!!")
	}

	resultProduct := db.FindLoginID(username, password)

	// URLへのアクセスに対してJSONを返す
	c.JSON(cnst.HttpStatusOK, resultProduct)
}

// クレジットカード情報を登録する
func FetchCreditInfoRegist(c *gin.Context) {
	cardnumber := c.PostForm("cardnumber")
	cardname := c.PostForm("cardname")
	cardmonth := c.PostForm("cardmonth")
	cardyear := c.PostForm("cardyear")
	cardcvv := c.PostForm("cardcvv")

	if len(cardnumber) == cnst.ZERO {
		log.Panic("Error nothing URL parameter!!")
	}

	cardmonthInt, _ := strconv.Atoi(cardmonth)
	cardyearInt, _ := strconv.Atoi(cardyear)

	resultProduct := db.AddCardInfo(cardnumber, cardname, cardmonthInt, cardyearInt, cardcvv)

	// URLへのアクセスに対してJSONを返す
	c.JSON(cnst.HttpStatusOK, resultProduct)
}
