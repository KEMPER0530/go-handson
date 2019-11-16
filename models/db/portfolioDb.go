package db

import (
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
func FindLoginID(username string) []entity.Login_info {
	login_info := []entity.Login_info{}

	db := open()
	// select
	db.First(&login_info, "username=?", username)
	close(db)

	return login_info
}

// work全件取得する
func FetchAllWorker() []entity.Work {
	work := []entity.Work{}

	db := open()
	db.Order("work_id asc").Find(&work)
	close(db)

	return work
}
