package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	// JobRunner
	"github.com/bamzi/jobrunner"

	// constクラス
	cnst "github.com/kemper0530/go-handson/common"
	// DBアクセス用モジュール
	db "github.com/kemper0530/go-handson/models/db"
)

// メールバッチ処理
func FetchMailSendSelect() {
	resultProduct := db.FetchMailSendSelect()
	fmt.Println("Run AmazonMail SES! %s", resultProduct)
}

// メールバッチステータスを返却する
func FetchMailBatchStatus(c *gin.Context) {
	c.JSON(http.StatusOK, jobrunner.StatusJson())
}

// FetchAllMembers は メンバー情報を取得する
func FetchAllMembers(c *gin.Context) {
	resultProducts := db.FetchAllMembers()

	// URLへのアクセスに対してJSONを返す
	c.JSON(http.StatusOK, resultProducts)
}

// work情報を取得する
func FetchAllWorker(c *gin.Context) {
	resultProducts := db.FetchAllWorker()

	// URLへのアクセスに対してJSONを返す
	c.JSON(http.StatusOK, resultProducts)
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
	c.JSON(http.StatusOK, resultProduct)
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
	c.JSON(http.StatusOK, resultProduct)
}

// お問合せ内容を登録する
func FetchSendMailRegist(c *gin.Context) {
	to_email := c.PostForm("to_email")
	name := c.PostForm("name")
	text := c.PostForm("text")
	from_email := c.PostForm("from_email")
	personal_name := c.PostForm("personal_name")

	if len(to_email) == cnst.ZERO &&
		len(name) == cnst.ZERO &&
		len(text) == cnst.ZERO {
		log.Panic("Error nothing URL parameter!!")
	}

	resultProduct := db.SendMailRegist(to_email, name, text, from_email, personal_name)

	// URLへのアクセスに対してJSONを返す
	c.JSON(http.StatusOK, resultProduct)
}
