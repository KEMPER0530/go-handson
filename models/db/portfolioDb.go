package db

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	// constクラス
	"github.com/google/uuid"
	cnst "github.com/kemper0530/go-handson/common"

	// エンティティ(データベースのテーブルの行に対応)
	entity "github.com/kemper0530/go-handson/models/entity"
)

// FindAllMembersはメンバー全件取得する
func FetchAllMembers() []entity.Testmember {
	testmember := []entity.Testmember{}

	db := open()
	db.Order("id asc").Find(&testmember)
	close(db)
	return testmember
}

// ログイン情報を取得する
func FindLoginID(username string, password string) entity.LoginRslt {
	login_info := []entity.Login_info{}
	loginrslt := entity.LoginRslt{}

	db := open()

	// select
	db.First(&login_info, "username=?", username)

	if len(login_info) == cnst.ONE {
		// verify
		errLogin := verify(login_info[0].Password, password)

		if errLogin == nil {
			fmt.Println("ok!")
			// ログイン成功
			loginrslt.Responce = cnst.JsonStatusOK
			loginrslt.Result = cnst.ONE
		} else {
			fmt.Println("err: ", errLogin)
			// ログイン失敗
			loginrslt.Responce = cnst.JsonStatusOK
			loginrslt.Result = cnst.ZERO
		}
	} else {
		fmt.Println("err no data: ")
		// ログイン失敗
		loginrslt.Responce = cnst.JsonStatusOK
		loginrslt.Result = cnst.ZERO
	}

	close(db)

	return loginrslt
}

// verify
func verify(hash, s string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(s))
}

// work全件取得する
func FetchAllWorker() []entity.Work {
	work := []entity.Work{}

	db := open()
	db.Order("work_id asc").Find(&work)
	close(db)

	return work
}

// カード情報を登録する
func AddCardInfo(cardnumber string, cardname string, cardmonth int, cardyear int, cardcvv string) entity.CrdRgstRslt {
	crdcardinfo := []entity.Crdcardinfo{}
	crdRgstRslt := entity.CrdRgstRslt{}

	// ハッシュ値の生成　セキュリティコードはbcryptで暗号化して登録
	hashCardcvv, err := bcrypt.GenerateFromPassword([]byte(cardcvv), bcrypt.DefaultCost)
	if err != nil {
		log.Panic("Error bcrypt.GenerateFromPassword!")
		crdRgstRslt.Responce = cnst.JsonStatusNG
		crdRgstRslt.Result = cnst.ZERO
		return crdRgstRslt
	}

	db := open()

	// select
	db.First(&crdcardinfo, "cardnumber=?", cardnumber)

	if len(crdcardinfo) == cnst.ONE {
		// 登録失敗
		crdRgstRslt.Responce = cnst.JsonStatusOK
		crdRgstRslt.Result = cnst.TWO
	} else {
		var crdcardinfoIns = entity.Crdcardinfo{
			Cardnumber: cardnumber,
			Cardname:   cardname,
			Cardmonth:  cardmonth,
			Cardyear:   cardyear,
			Cardcvv:    string(hashCardcvv),
		}
		// insert
		db.Create(&crdcardinfoIns)
		crdRgstRslt.Responce = cnst.JsonStatusOK
		crdRgstRslt.Result = cnst.ONE
	}

	close(db)

	return crdRgstRslt
}

// お問合せ内容を登録する
func SendMailRegist(to_email string, name string, text string, from_email string, personal_name string) entity.SendMailRslt {

	mail_send_rslt := []entity.Mail_send_rslt{}
	mst_ssmlknr := []entity.Mst_ssmlknr{}
	sendMailRslt := entity.SendMailRslt{}

	db := open()

	// 送信連番の取得
	count := cnst.ZERO
	sendno1 := cnst.ZERO
	db.Find(&mail_send_rslt).Count(&count)
	sendno1 = count + 1
	sendno2 := sendno1 + 1

	// テナントIDの定義
	tnntid := cnst.TNNTID

	// msg_id1の生成
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		// 登録失敗
		sendMailRslt.Responce = cnst.JsonStatusOK
		sendMailRslt.Result = cnst.TWO
		return sendMailRslt
	}
	msgid1 := u.String()

	// msg_id2の生成
	u2, err2 := uuid.NewRandom()
	if err2 != nil {
		fmt.Println(err2)
		// 登録失敗
		sendMailRslt.Responce = cnst.JsonStatusOK
		sendMailRslt.Result = cnst.TWO
		return sendMailRslt
	}
	msgid2 := u2.String()

	// 送信管理マスタの取得
	db.First(&mst_ssmlknr)
	subject := mst_ssmlknr[0].Subject
	body := mst_ssmlknr[0].Body
	replytitle := mst_ssmlknr[0].Replytitle
	toreply := mst_ssmlknr[0].Toreply

	// 文字列の置き換え　$1　→　登録名
	_body := strings.Replace(body, "$1", name, -1)

	// 日本時間へ変換
	jst, _ := time.LoadLocation("Asia/Tokyo")
	_time := time.Now().In(jst)

	// insert メール送信情報(顧客用)
	var mail_send_infIns = entity.Mail_send_inf{
		Msg_id:        msgid1,
		From_email:    from_email,
		To_email:      to_email,
		Subject:       subject,
		Body:          _body,
		Personal_name: personal_name,
	}

	// insert メール送信結果情報(顧客用)
	var mail_send_rsltIns = entity.Mail_send_rslt{
		Send_no:         sendno1,
		Msg_id:          msgid1,
		Tnnt_id:         tnntid,
		Target_sys_type: strconv.Itoa(cnst.ONE),
		Status:          strconv.Itoa(cnst.ZERO),
		Server_id:       cnst.SERVID,
		Priority:        cnst.ONE,
		Send_reg_at:     _time,
		Queue_remove:    strconv.Itoa(cnst.ZERO),
		Updated_at:      _time,
	}

	// insert
	db.Create(&mail_send_infIns)
	db.Create(&mail_send_rsltIns)

	// insert メール送信情報(送信者用)
	mail_send_infIns = entity.Mail_send_inf{
		Msg_id:        msgid2,
		From_email:    from_email,
		To_email:      toreply,
		Subject:       replytitle,
		Body:          text,
		Personal_name: personal_name,
	}

	// insert メール送信結果情報(送信者用)
	mail_send_rsltIns = entity.Mail_send_rslt{
		Send_no:         sendno2,
		Msg_id:          msgid2,
		Tnnt_id:         tnntid,
		Target_sys_type: strconv.Itoa(cnst.ONE),
		Status:          strconv.Itoa(cnst.ZERO),
		Server_id:       cnst.SERVID,
		Priority:        cnst.ONE,
		Send_reg_at:     _time,
		Queue_remove:    strconv.Itoa(cnst.ZERO),
		Updated_at:      _time,
	}

	// insert
	db.Create(&mail_send_infIns)
	db.Create(&mail_send_rsltIns)

	sendMailRslt.Responce = cnst.JsonStatusOK
	sendMailRslt.Result = cnst.ONE

	close(db)

	return sendMailRslt
}
