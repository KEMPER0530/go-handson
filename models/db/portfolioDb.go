package db

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	// constクラス
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

	// ハッシュ値の生成
	// hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// if err != nil {
	// 	log.Panic("Error bcrypt.GenerateFromPassword!")
	// 	loginrslt.Responce = 500
	// 	loginrslt.Result = 0
	// 	return loginrslt
	// }

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
