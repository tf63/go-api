/*
# UserのRepository層
  - Create, Read操作を定義
*/
package repository

import (
	"time"

	"github.com/tf63/go_api/internal/entity"
	"gorm.io/gorm"
)

/*
UserRepository Interface
  - Mockするので抽象化
*/
type UserRepository interface {
	CreateUser(input entity.NewUser) (userId int, err error)
	ReadUser(userId int) (user entity.User, err error)
	ReadUsers() (users []entity.User, err error)
	UpdateUser(input entity.NewUser, userId int) (err error)
	DeleteUser(userId int) (err error)
}

// UserRepository 構造体
type userRepository struct {
	db gorm.DB
}

// Make sure we conform to ServerInterface
var _ UserRepository = (*userRepository)(nil)

// インスタンスの取得?
func NewUserRepository(db gorm.DB) UserRepository {
	return &userRepository{db}
}

/*
Create: Todoを作成する
  - input:
  - Name string `json:"name"`
  - return:
  - None
  - Error:
  - STATUS_SERVICE_UNAVAILABLE (503)
*/
func (ur *userRepository) CreateUser(input entity.NewUser) (userId int, err error) {

	// inputを取得
	if input.Name == nil {
		err = entity.STATUS_SERVICE_UNAVAILABLE
		return
	}

	name := *input.Name

	query := `
	INSERT INTO users (name, created_at, updated_at)
	VALUES (?, ?, ?)
	`

	args := []interface{}{
		name,
		time.Now(),
		time.Now(),
	}

	// レコードの作成
	result := ur.db.Exec(query, args...)

	if result.Error != nil {
		err = entity.STATUS_SERVICE_UNAVAILABLE
		return
	}

	result = ur.db.Raw(`SELECT id FROM users ORDER BY id DESC LIMIT 1`).Scan(&userId)
	if result.Error != nil {
		err = entity.STATUS_SERVICE_UNAVAILABLE
		return
	}

	return
}

/*
Read: UserをuserIDで指定して1件取得する (nameからuserID等を取得すべきかもしれない)
  - input:
  - UserID string `json:"userId"`
  - return:
  - User 1件
  - Error:
  - STATUS_SERVICE_UNAVAILABLE (503)
*/
func (ur *userRepository) ReadUser(userId int) (user entity.User, err error) {

	// userIDからレコードを取得
	record := entity.User{}

	query := "SELECT * FROM users WHERE id = ?"
	args := []interface{}{uint(userId)}

	// レコードを割り当てる
	result := ur.db.Raw(query, args...).Scan(&record)

	// レコードが存在しなかったら (汚い)
	if record.ID == 0 {
		err = entity.STATUS_NOT_FOUND
		return
	}

	if result.Error == gorm.ErrRecordNotFound {
		// gormのエラーの種類でユーザーが存在するかどうかわかる
		// 意味ないが後学のための残しておく
		err = entity.STATUS_SERVICE_UNAVAILABLE
		return
	} else if result.Error != nil {
		err = entity.STATUS_SERVICE_UNAVAILABLE
		return
	}

	user = record
	return
}

/*
Read: Userをlimit件取得する (limitは入力として与えるべきかもしれない)
  - input: None
  - return:
  - User 1件
  - Error:
  - STATUS_SERVICE_UNAVAILABLE (503)
*/
func (ur *userRepository) ReadUsers() (users []entity.User, err error) {

	// レコードをlimit件取得
	record := []entity.User{}

	limit := 500

	query := "SELECT * FROM users LIMIT ?"
	args := []interface{}{uint(limit)}

	// レコードを割り当てる
	result := ur.db.Raw(query, args...).Scan(&record)

	if result.Error == gorm.ErrRecordNotFound {
		// gormのエラーの種類でユーザーが存在するかどうかわかる
		// 意味ないが後学のための残しておく
		err = entity.STATUS_SERVICE_UNAVAILABLE
		return
	} else if result.Error != nil {
		err = entity.STATUS_SERVICE_UNAVAILABLE
		return
	}

	users = record
	return
}

/*
Update: userをuserIdで指定して更新する
  - input:
  - Name *string `json:"name,omitempty"`
  - userId int
  - return:
  - None
  - Error:
  - STATUS_NOT_FOUND (404)
  - STATUS_SERVICE_UNAVAILABLE (503)
*/
func (ur *userRepository) UpdateUser(input entity.NewUser, userId int) (err error) {

	// レコードの更新
	query := `UPDATE users SET `
	args := []interface{}{}

	if input.Name != nil {
		query += "name = ?, "
		args = append(args, *input.Name)
	}

	query += "updated_at = ? "
	args = append(args, time.Now())

	query += "WHERE id = ?"
	args = append(args, userId)

	// Updateの実行
	result := ur.db.Exec(query, args...)

	// エラーハンドリング
	if result.Error != nil {
		err = entity.STATUS_SERVICE_UNAVAILABLE // 503
		return
	} else if result.RowsAffected == 0 {
		// 更新の結果，影響がなかったらtodoIdが存在しないと考える
		err = entity.STATUS_NOT_FOUND // 404
		return
	}

	return
}

/*
Delete: userをuserIdで指定して削除する
  - input:
  - userId int
  - return:
  - None
  - Error:
  - STATUS_NOT_FOUND (404)
  - STATUS_SERVICE_UNAVAILABLE (503)
*/
func (ur *userRepository) DeleteUser(userId int) (err error) {

	// userIdに対応するレコードを削除する
	query := `DELETE FROM users WHERE id = ?`
	args := []interface{}{userId}

	result := ur.db.Exec(query, args...)

	// エラーハンドリング
	if result.Error != nil {
		return entity.STATUS_SERVICE_UNAVAILABLE // 503
	} else if result.RowsAffected == 0 {
		// 削除の結果，影響がなかったらtodoIdが存在しないと考える
		return entity.STATUS_NOT_FOUND // 404
	}

	return nil
}
