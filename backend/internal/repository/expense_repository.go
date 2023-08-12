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

	"github.com/tf63/go_api/api/rest"
	"github.com/tf63/go_api/external"
	"github.com/tf63/go_api/internal/entity"
	"gorm.io/gorm"
)

type ExpenseRepository interface {
	CreateExpense(input rest.NewExpense) (expense_id int, err error)
	ReadExpense(input rest.FindUser, expense_id int) (expense entity.Expense, err error)
	ReadExpenses(input rest.FindUser) (expenses []entity.Expense, err error)
	UpdateExpense(input rest.NewExpense, expense_id int) (err error)
	DeleteExpense(input rest.FindUser, expense_id int) (err error)
}

type expenseRepository struct {
	db gorm.DB
}

func NewExpenseRepository(db gorm.DB) ExpenseRepository {
	return &expenseRepository{db}
}

func GetGroupId(expense_id int) string {
	div_size := external.GetDivSize()
	return strconv.Itoa((expense_id % div_size) + 1)
}

/*
Create: Expenseを作成する
  - input:
  - Price  *int    `json:"price,omitempty"`
  - Title  *string `json:"title,omitempty"`
  - ExpenseId *int    `json:"expense_id,omitempty"`
  - return:
  - None
  - Error:
  - STATUS_SERVICE_UNAVAILABLE (503)
*/
func (er *expenseRepository) CreateExpense(input rest.NewExpense) (expense_id int, err error) {

	// inputを取得
	if input.Title == nil || input.Price == nil || input.UserId == nil {
		err = entity.STATUS_SERVICE_UNAVAILABLE
		return
	}

	title := *input.Title
	price := *input.Price
	user_id := *input.UserId

	// user_idでレコードを絞る
	user_group := GetGroupId(user_id)

	query := `
	INSERT INTO expenses_` + user_group + ` (title, price, expense_id, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?)
	`

	args := []interface{}{
		title,
		price,
		uint(user_id),
		time.Now(),
		time.Now(),
	}

	// レコードの作成
	result := er.db.Exec(query, args...)
	if result.Error != nil {
		err = entity.STATUS_SERVICE_UNAVAILABLE
		return
	}

	result = er.db.Raw(`SELECT id FROM expenses_` + user_group + ` ORDER BY id DESC LIMIT 1`).Scan(&expense_id)
	if result.Error != nil {
		err = entity.STATUS_SERVICE_UNAVAILABLE
		return
	}

	return
}

func (er *expenseRepository) ReadExpense(input rest.FindUser, expense_id int) (expense entity.Expense, err error) {

	if input.UserId == nil {
		err = entity.STATUS_SERVICE_UNAVAILABLE
		return
	}

	user_id := *input.UserId

	// user_idでレコードを絞る
	user_group := GetGroupId(user_id)

	// expense_idからレコードを取得
	record := entity.Expense{}

	query := `SELECT * FROM expenses_` + user_group + ` WHERE id = ? and AND user_id = ?`
	args := []interface{}{uint(expense_id), uint(user_id)}

	// レコードを割り当てる
	result := er.db.Raw(query, args...).Scan(&record)

	if result.Error == gorm.ErrRecordNotFound {
		// gormのエラーの種類で存在するかどうかわかる
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

func (er *expenseRepository) ReadExpenses(input rest.FindUser) (expenses []entity.Expense, err error) {

	if input.UserId == nil {
		err = entity.STATUS_SERVICE_UNAVAILABLE
		return
	}

	// レコードをlimit件取得
	record := []entity.Expense{}

	user_id := *input.UserId

	// user_idでレコードを絞る
	user_group := GetGroupId(user_id)

	limit := 500

	query := `SELECT * FROM expenses_` + user_group + ` LIMIT ?`
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

func (er *expenseRepository) UpdateExpense(input rest.NewExpense, expense_id int) (err error) {

	if input.UserId == nil {
		err = entity.STATUS_SERVICE_UNAVAILABLE
		return
	}

	user_id := *input.UserId

	// user_idでレコードを絞る
	user_group := GetGroupId(user_id)

	// レコードの更新
	query := `UPDATE expenses_` + user_group + ` SET `
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
	args = append(args, expense_id)

	// 他のuserのtodoを更新できないようにする
	query += `AND user_id = ?`
	args = append(args, user_id)

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

func (er *expenseRepository) DeleteExpense(input rest.FindUser, expense_id int) (err error) {

	if input.UserId == nil {
		err = entity.STATUS_SERVICE_UNAVAILABLE
		return
	}

	user_id := *input.UserId

	// user_idでレコードを絞る
	user_group := GetGroupId(user_id)

	// expense_idに対応するレコードを削除する
	query := `DELETE FROM expenses_` + user_group + ` WHERE id = ? AND user_id = ?`
	args := []interface{}{expense_id, user_id}

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
