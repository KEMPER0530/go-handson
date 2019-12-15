package db

import (
	// フォーマットI/O

	"fmt"
	"log"
	"os"

	// Go言語のORM

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// DB接続する
func open() *gorm.DB {

	// 環境変数から値を取得
	dbms := os.Getenv("DBMS")
	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASS")
	dbprotocol := os.Getenv("DB_PROTOCOL")
	dbname := os.Getenv("DB_NAME")
	connect := dbuser + ":" + dbpass + "@" + dbprotocol + "/" + dbname

	//mysqlへ接続
	db, err := gorm.Open(dbms, connect)

	//接続でエラーが発生した場合の処理
	if err != nil {
		panic(err.Error())
	}

	// DBエンジンを「InnoDB」に設定
	db.Set("gorm:table_options", "ENGINE=InnoDB")

	// 詳細なログを表示
	db.LogMode(true)

	// 登録するテーブル名を単数形にする（デフォルトは複数形）
	db.SingularTable(true)

	fmt.Println("db connected: ", &db)
	return db
}

func close(db *gorm.DB) {
	defer db.Close()
}

func awsEnvload(_num string) string {
	err := godotenv.Load(fmt.Sprintf("config/production.env"))
	if err != nil {
		log.Panic("Error loading .env file")
	}

	envString := os.Getenv(_num)

	return envString
}
