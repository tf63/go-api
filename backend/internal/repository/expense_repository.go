/*
# TodoのRepository層
  - CRUD操作を定義
  - (SelectではなくReadのほうが良かった気がする)
  - (SelectTodoはいらない気がする)

参考 https://zenn.dev/maruware/scraps/1a71e4664b1fae
*/
package repository

import (
	"strconv"
	"time"

	"github.com/tf63/go_api/external"
	"github.com/tf63/go_api/internal/entity"
	"gorm.io/gorm"
)

type ExpenseRepository interface {
	CreateExpense(input entity.NewExpense) (expenseId int, err error)
	ReadExpense(input entity.FindUser, expenseId int) (expense entity.Expense, err error)
	ReadExpenses(input entity.FindUser) (expenses []entity.Expense, err error)
	UpdateExpense(input entity.NewExpense, expenseId int) (err error)
	DeleteExpense(input entity.FindUser, expenseId int) (err error)
}

type expenseRepository struct {
	db gorm.DB
}

func NewExpenseRepository(db gorm.DB) ExpenseRepository {
	return &expenseRepository{db}
}

func GetGroupId(userId uint) string {
	divSize := external.GetDivSize()
	return strconv.Itoa((int(userId) % divSize) + 1)
}

/*
Create: Expenseを作成する
  - input:
  - Price  *int    `json:"price,omitempty"`
  - Title  *string `json:"title,omitempty"`
  - ExpenseId *int    `json:"expenseId,omitempty"`
  - return:
  - None
  - Error:
  - STATUS_SERVICE_UNAVAILABLE (503)
*/
func (er *expenseRepository) CreateExpense(input entity.NewExpense) (expenseId int, err error) {

	// inputを取得
	if input.Title == nil || input.Price == nil {
		err = entity.STATUS_SERVICE_UNAVAILABLE
		return
	}

	title := *input.Title
	price := *input.Price
	userId := input.UserID

	// userIdでレコードを絞る
	userGroup := GetGroupId(userId)

	query := `
	INSERT INTO expenses_` + userGroup + ` (title, price, user_id, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?)
	`

	args := []interface{}{
		title,
		price,
		uint(userId),
		time.Now(),
		time.Now(),
	}

	// レコードの作成
	result := er.db.Exec(query, args...)
	if result.Error != nil {
		err = entity.STATUS_SERVICE_UNAVAILABLE
		return
	}

	result = er.db.Raw(`SELECT id FROM expenses_` + userGroup + ` ORDER BY id DESC LIMIT 1`).Scan(&expenseId)
	if result.Error != nil {
		err = entity.STATUS_SERVICE_UNAVAILABLE
		return
	}

	return
}

func (er *expenseRepository) ReadExpense(input entity.FindUser, expenseId int) (expense entity.Expense, err error) {

	// if input.UserId == nil {
	// 	err = entity.STATUS_SERVICE_UNAVAILABLE
	// 	return
	// }

	userId := input.ID

	// userIdでレコードを絞る
	userGroup := GetGroupId(userId)

	// expenseIdからレコードを取得
	record := entity.Expense{}

	query := `SELECT * FROM expenses_` + userGroup + ` WHERE id = ? AND user_id = ?`
	args := []interface{}{uint(expenseId), uint(userId)}

	// レコードを割り当てる
	result := er.db.Raw(query, args...).Scan(&record)

	// レコードが存在しなかったら (汚い)
	if record.ID == 0 {
		err = entity.STATUS_NOT_FOUND
		return
	}

	if result.Error == gorm.ErrRecordNotFound {
		// gormのエラーの種類で存在するかどうかわかる -> db.Rawでは判定してくれない?
		// 意味ないが後学のための残しておく
		err = entity.STATUS_SERVICE_UNAVAILABLE
		return
	} else if result.Error != nil {
		err = entity.STATUS_SERVICE_UNAVAILABLE
		return
	}

	expense = record
	return
}

func (er *expenseRepository) ReadExpenses(input entity.FindUser) (expenses []entity.Expense, err error) {

	// if input.UserId == nil {
	// 	err = entity.STATUS_SERVICE_UNAVAILABLE
	// 	return
	// }

	// レコードをlimit件取得
	record := []entity.Expense{}

	userId := input.ID

	// userIdでレコードを絞る
	userGroup := GetGroupId(userId)

	limit := 500

	query := `SELECT * FROM expenses_` + userGroup + ` LIMIT ?`
	args := []interface{}{uint(limit)}

	// レコードを割り当てる
	result := er.db.Raw(query, args...).Scan(&record)

	if result.Error == gorm.ErrRecordNotFound {
		// gormのエラーの種類でユーザーが存在するかどうかわかる
		// 意味ないが後学のための残しておく
		err = entity.STATUS_SERVICE_UNAVAILABLE
		return
	} else if result.Error != nil {
		err = entity.STATUS_SERVICE_UNAVAILABLE
		return
	}

	expenses = record
	return
}

func (er *expenseRepository) UpdateExpense(input entity.NewExpense, expenseId int) (err error) {

	// if input.UserId == nil {
	// 	err = entity.STATUS_SERVICE_UNAVAILABLE
	// 	return
	// }

	userId := input.UserID

	// userIdでレコードを絞る
	userGroup := GetGroupId(userId)

	// レコードの更新
	query := `UPDATE expenses_` + userGroup + ` SET `
	args := []interface{}{}

	if input.Title != nil {
		query += "title = ?, "
		args = append(args, *input.Title)
	}

	if input.Price != nil {
		query += "price = ?, "
		args = append(args, *input.Price)
	}

	query += "updated_at = ? "
	args = append(args, time.Now())

	query += "WHERE id = ? "
	args = append(args, expenseId)

	// 他のuserのtodoを更新できないようにする
	query += `AND user_id = ?`
	args = append(args, userId)

	// Updateの実行
	result := er.db.Exec(query, args...)

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

func (er *expenseRepository) DeleteExpense(input entity.FindUser, expenseId int) (err error) {

	userId := input.ID

	// userIdでレコードを絞る
	userGroup := GetGroupId(userId)

	// expenseIdに対応するレコードを削除する
	query := `DELETE FROM expenses_` + userGroup + ` WHERE id = ? AND user_id = ?`
	args := []interface{}{expenseId, userId}

	result := er.db.Exec(query, args...)

	// エラーハンドリング
	if result.Error != nil {
		return entity.STATUS_SERVICE_UNAVAILABLE // 503
	} else if result.RowsAffected == 0 {
		// 削除の結果，影響がなかったらtodoIdが存在しないと考える
		return entity.STATUS_NOT_FOUND // 404
	}

	return
}
