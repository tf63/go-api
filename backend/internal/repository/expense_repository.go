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
}

type expenseRepository struct {
	db gorm.DB
}

func NewExpenseRepository(db gorm.DB) ExpenseRepository {
	return &expenseRepository{db}
}

func GetGroupId(user_id int) string {
	div_size := external.GetDivSize()
	return strconv.Itoa((user_id % div_size) + 1)
}

/*
Create: Expenseを作成する
  - input:
  - Price  *int    `json:"price,omitempty"`
  - Title  *string `json:"title,omitempty"`
  - UserId *int    `json:"user_id,omitempty"`
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

	// userIDでレコードを絞る
	user_group := GetGroupId(user_id)

	query := `
	INSERT INTO expenses_` + user_group + ` (title, price, user_id, created_at, updated_at)
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
