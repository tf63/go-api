/*
# データベースへの接続を定義する
  - GORMで接続する
  - GORMでマイグレーションを行う

参考 https://zenn.dev/maruware/scraps/1a71e4664b1fae
公式 https://gorm.io/ja_JP/docs/connecting_to_the_database.html
*/
package external

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/tf63/go_api/internal/entity"
)

const div_size = 16

// GORMでDBに接続し，GORMのDBオブジェクトを返す
func ConnectDatabase(isTest bool) (*gorm.DB, error) {
	dsn := ""
	// DBへの接続に必要な情報
	if isTest {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
			os.Getenv("POSTGRES_HOST_TEST"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_NAME"),
			// os.Getenv("POSTGRES_PORT"),
			"5432",
		)
	} else {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_NAME"),
			os.Getenv("POSTGRES_PORT"),
		)
	}

	// DBへ接続
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// テスト時はテーブルを初期化する
	if isTest {
		db.Migrator().DropTable(&entity.User{})
		for i := 1; i <= div_size; i++ {
			table_name := "expenses_" + strconv.Itoa(i)
			db.Migrator().DropTable(table_name)
		}
	}

	// マイグレーションを実行
	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		return nil, err
	}

	// expenseテーブルを複数作成する
	for i := 1; i <= div_size; i++ {
		table_name := "expenses_" + strconv.Itoa(i)
		if !db.Migrator().HasTable(table_name) {
			err = db.Table(table_name).Migrator().CreateTable(&entity.Expense{})

			if err != nil {
				return nil, err
			}
		}
	}

	return db, nil
}

func GetDivSize() int {
	return div_size
}
