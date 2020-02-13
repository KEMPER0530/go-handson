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
func FindLoginID(username string, password string) entity.Rslt {
	login_info := []entity.Login_info{}
	Rslt := entity.Rslt{}

	db := open()

	// select
	db.First(&login_info, "username=?", username)

	if len(login_info) == cnst.ONE {
		// verify
		errLogin := verify(login_info[0].Password, password)

		if errLogin == nil {
			fmt.Println("Login success!")
			// ログイン成功
			Rslt.Responce = cnst.JsonStatusOK
			Rslt.Result = cnst.ONE
			Rslt.Name = login_info[0].Name
		} else {
			fmt.Println("Login error: ", errLogin)
			// ログイン失敗
			Rslt.Responce = cnst.JsonStatusOK
			Rslt.Result = cnst.ZERO
		}
	} else {
		fmt.Println("Login error no data: ")
		// ログイン失敗
		Rslt.Responce = cnst.JsonStatusOK
		Rslt.Result = cnst.ZERO
	}

	close(db)

	return Rslt
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
func AddCardInfo(cardnumber string, cardname string, cardmonth int, cardyear int, cardcvv string) entity.Rslt {
	crdcardinfo := []entity.Crdcardinfo{}
	Rslt := entity.Rslt{}

	// ハッシュ値の生成　セキュリティコードはbcryptで暗号化して登録
	hashCardcvv, err := bcrypt.GenerateFromPassword([]byte(cardcvv), bcrypt.DefaultCost)
	if err != nil {
		log.Panic("Error bcrypt.GenerateFromPassword!")
		Rslt.Responce = cnst.JsonStatusNG
		Rslt.Result = cnst.ZERO
		return Rslt
	}

	db := open()

	// select
	db.First(&crdcardinfo, "cardnumber=?", cardnumber)

	if len(crdcardinfo) == cnst.ONE {
		// 登録失敗
		Rslt.Responce = cnst.JsonStatusOK
		Rslt.Result = cnst.TWO
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
		Rslt.Responce = cnst.JsonStatusOK
		Rslt.Result = cnst.ONE
	}

	close(db)

	return Rslt
}

// メール送信結果テーブルを設定する
func SetMailSendRslt() entity.Mail_send_rslt {

	mail_send_rslt := []entity.Mail_send_rslt{}

	db := open()

	// 送信連番の取得
	count := cnst.ZERO
	sendno := cnst.ZERO

	db.Find(&mail_send_rslt).Count(&count)
	sendno = count + 1

	// テナントIDの定義
	tnntid := cnst.TNNTID

	// msg_idの生成
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		// 登録失敗
		return mail_send_rslt[0]
	}
	msgid := u.String()

	// 日本時間へ変換
	jst, _ := time.LoadLocation("Asia/Tokyo")
	_time := time.Now().In(jst)

	// insert メール送信結果情報(顧客用)
	mail_send_rsltIns := entity.Mail_send_rslt{
		Send_no:         sendno,
		Msg_id:          msgid,
		Tnnt_id:         tnntid,
		Target_sys_type: strconv.Itoa(cnst.ONE),
		Status:          strconv.Itoa(cnst.ZERO),
		Server_id:       cnst.SERVID,
		Priority:        cnst.ONE,
		Send_reg_at:     _time,
		Queue_remove:    strconv.Itoa(cnst.ZERO),
		Updated_at:      _time,
	}

	close(db)

	return mail_send_rsltIns

}

// メール送信情報テーブルを設定
func SetMailSendInf2C(to_email string, name string, text string, from_email string, personal_name string, msgid string, id int) entity.Mail_send_inf {

	mst_ssmlknr := []entity.Mst_ssmlknr{}

	db := open()

	// 送信管理マスタの取得
	db.Where("id = ?", id).First(&mst_ssmlknr)
	subject := mst_ssmlknr[0].Subject
	body := mst_ssmlknr[0].Body

	// 文字列の置き換え　$1　→　登録名
	_body := strings.Replace(body, "$1", name, -1)

	// insert メール送信情報(顧客用)
	mail_send_infIns := entity.Mail_send_inf{
		Msg_id:        msgid,
		From_email:    from_email,
		To_email:      to_email,
		Subject:       subject,
		Body:          _body,
		Personal_name: personal_name,
	}

	close(db)

	return mail_send_infIns
}

// メール送信情報テーブルを設定
func SetMailSendInf2Y(to_email string, name string, text string, from_email string, personal_name string, msgid string, id int) entity.Mail_send_inf {

	mst_ssmlknr := []entity.Mst_ssmlknr{}

	db := open()

	// 送信管理マスタの取得
	db.Where("id = ?", id).First(&mst_ssmlknr)
	replytitle := mst_ssmlknr[0].Replytitle
	toreply := mst_ssmlknr[0].Toreply

	// insert メール送信情報(送信者用)
	mail_send_infIns := entity.Mail_send_inf{
		Msg_id:        msgid,
		From_email:    from_email,
		To_email:      toreply,
		Subject:       replytitle,
		Body:          text,
		Personal_name: personal_name,
	}

	close(db)

	return mail_send_infIns
}

// お問合せ内容をメール送信情報テーブル、結果テーブルへ登録する
func SetMailRegist(sendInf *entity.Mail_send_inf, sendRslt *entity.Mail_send_rslt) entity.Rslt {

	Rslt := entity.Rslt{}

	db := open()

	// insert
	db.Create(&sendInf)
	db.Create(&sendRslt)

	Rslt.Responce = cnst.JsonStatusOK
	Rslt.Result = cnst.ONE

	close(db)

	return Rslt
}

// お問合せ内容をメール送信情報テーブル、結果テーブルへ登録する
func SendMailRegist(to_email string, name string, text string, from_email string, personal_name string) entity.Rslt {

	mail_send_rslt := []entity.Mail_send_rslt{}
	mst_ssmlknr := []entity.Mst_ssmlknr{}
	Rslt := entity.Rslt{}

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
		Rslt.Responce = cnst.JsonStatusOK
		Rslt.Result = cnst.TWO
		return Rslt
	}
	msgid1 := u.String()

	// msg_id2の生成
	u2, err2 := uuid.NewRandom()
	if err2 != nil {
		fmt.Println(err2)
		// 登録失敗
		Rslt.Responce = cnst.JsonStatusOK
		Rslt.Result = cnst.TWO
		return Rslt
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

	Rslt.Responce = cnst.JsonStatusOK
	Rslt.Result = cnst.ONE

	close(db)

	return Rslt
}

// profile全件取得する
func FetchProfileInfo() []entity.Profile {
	profile := []entity.Profile{}

	db := open()
	db.Order("id asc").Find(&profile)
	close(db)

	return profile
}

// アカウント情報を登録する
func RegistLoginID(username string, password string, name string) entity.Rslt {
	login_info := []entity.Login_info{}
	Rslt := entity.Rslt{}

	db := open()

	// 登録情報の確認
	db.First(&login_info, "username=?", username)

	if len(login_info) == cnst.ONE {
		Rslt.Responce = cnst.JsonStatusOK
		Rslt.Result = cnst.ZERO
		return Rslt
	}

	// ハッシュ値の生成　パスワードはbcryptで暗号化して登録
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Panic("Error bcrypt.GenerateFromPassword!")
		Rslt.Responce = cnst.JsonStatusNG
		Rslt.Result = cnst.ZERO
		return Rslt
	}

	// insert ログイン情報
	var login_infIns = entity.Login_info{
		Username: username,
		Password: string(hashPassword),
		Name:     name,
	}

	// insert
	db.Create(&login_infIns)

	Rslt.Responce = cnst.JsonStatusOK
	Rslt.Result = cnst.ONE

	close(db)

	return Rslt
}
