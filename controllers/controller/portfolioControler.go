package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	// JobRunner
	"github.com/bamzi/jobrunner"

	// authクラス
	authcnfg "github.com/kemper0530/go-handson/config"
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
	resultStatus, errMsg := authcnfg.AuthFirebase(c)
	if resultStatus == cnst.JsonStatusNG {
		c.JSON(http.StatusBadRequest, errMsg)
	} else {
		resultProducts := db.FetchAllMembers()
		// URLへのアクセスに対してJSONを返す
		c.JSON(http.StatusOK, resultProducts)
	}
}

// work情報を取得する
func FetchAllWorker(c *gin.Context) {
	resultStatus, errMsg := authcnfg.AuthFirebase(c)
	if resultStatus == cnst.JsonStatusNG {
		c.JSON(http.StatusBadRequest, errMsg)
	} else {
		resultProducts := db.FetchAllWorker()
		// URLへのアクセスに対してJSONを返す
		c.JSON(http.StatusOK, resultProducts)
	}
}

// FetchLoginInfo は 指定したIDのパスワードを取得する
func FetchLoginInfo(c *gin.Context) {
	resultStatus, errMsg := authcnfg.AuthFirebase(c)
	if resultStatus == cnst.JsonStatusNG {
		c.JSON(http.StatusBadRequest, errMsg)
	} else {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if len(username) == cnst.ZERO || len(password) == cnst.ZERO {
			log.Panic("Error nothing URL parameter!!")
		}

		resultProduct := db.FindLoginID(username, password)

		// URLへのアクセスに対してJSONを返す
		c.JSON(http.StatusOK, resultProduct)
	}
}

// クレジットカード情報を登録する
func FetchCreditInfoRegist(c *gin.Context) {
	resultStatus, errMsg := authcnfg.AuthFirebase(c)
	if resultStatus == cnst.JsonStatusNG {
		c.JSON(http.StatusBadRequest, errMsg)
	} else {
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
}

// お問合せ内容を登録する
func FetchSendMailRegist(c *gin.Context) {
	resultStatus, errMsg := authcnfg.AuthFirebase(c)
	if resultStatus == cnst.JsonStatusNG {
		c.JSON(http.StatusBadRequest, errMsg)
	} else {
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

		// 顧客向けのメール情報を設定する
		Mail_send_rslt := db.SetMailSendRslt()
		Mail_send_inf := db.SetMailSendInf2C(to_email, name, text, from_email, personal_name, Mail_send_rslt.Msg_id, cnst.ONE)
		resultProduct := db.SetMailRegist(&Mail_send_inf, &Mail_send_rslt)

		// メールバッチ処理を直接コールする
		FetchMailSendSelect()

		// 管理者向けのメール情報を設定する
		Mail_send_rslt = db.SetMailSendRslt()
		Mail_send_inf = db.SetMailSendInf2Y(to_email, name, text, from_email, personal_name, Mail_send_rslt.Msg_id, cnst.ONE)
		resultProduct = db.SetMailRegist(&Mail_send_inf, &Mail_send_rslt)

		// URLへのアクセスに対してJSONを返す
		c.JSON(http.StatusOK, resultProduct)

		// メールバッチ処理を直接コールする
		FetchMailSendSelect()
	}
}

// profile情報を取得する
func FetchProfileInfo(c *gin.Context) {
	resultStatus, errMsg := authcnfg.AuthFirebase(c)
	if resultStatus == cnst.JsonStatusNG {
		c.JSON(http.StatusBadRequest, errMsg)
	} else {
		resultProducts := db.FetchProfileInfo()

		// URLへのアクセスに対してJSONを返す
		c.JSON(http.StatusOK, resultProducts)
	}
}

// FetchRegistAcount は アカウントの登録を実施する
func FetchRegistAccount(c *gin.Context) {
	resultStatus, errMsg := authcnfg.AuthFirebase(c)
	if resultStatus == cnst.JsonStatusNG {
		c.JSON(http.StatusBadRequest, errMsg)
	} else {
		username := c.PostForm("email")
		password := c.PostForm("password")
		name := c.PostForm("name")

		if len(username) == cnst.ZERO || len(password) == cnst.ZERO {
			log.Panic("Error nothing URL parameter!!")
		}

		resultProduct := db.RegistLoginID(username, password, name)

		// URLへのアクセスに対してJSONを返す
		c.JSON(http.StatusOK, resultProduct)
	}
}

// FetchRegistAcountMail は 送信先へのメール情報を登録する
func FetchRegistAccountMail(c *gin.Context) {
	resultStatus, errMsg := authcnfg.AuthFirebase(c)
	if resultStatus == cnst.JsonStatusNG {
		c.JSON(http.StatusBadRequest, errMsg)
	} else {
		to_email := c.PostForm("to_email")
		name := c.PostForm("name")
		text := c.PostForm("text")
		from_email := c.PostForm("from_email")
		personal_name := c.PostForm("personal_name")

		if len(to_email) == cnst.ZERO {
			log.Panic("Error nothing URL parameter!!")
		}

		// 顧客向けのメール情報を設定する
		Mail_send_rslt := db.SetMailSendRslt()
		Mail_send_inf := db.SetMailSendInf2C(to_email, name, text, from_email, personal_name, Mail_send_rslt.Msg_id, cnst.TWO)
		resultProduct := db.SetMailRegist(&Mail_send_inf, &Mail_send_rslt)

		// メールバッチ処理を直接コールする
		FetchMailSendSelect()

		// URLへのアクセスに対してJSONを返す
		c.JSON(http.StatusOK, resultProduct)
	}
}
