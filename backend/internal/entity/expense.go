/*
# ExpenseのEntity層
  - GORMでマイグレーションを行う

参考 https://zenn.dev/maruware/scraps/1a71e4664b1fae
*/

package entity

import (
	"gorm.io/gorm"
)

/*
ExpenseEntity
  - Text: タスクのタイトル
  - Price: 価格
  - UserID: ユーザーID (FK)
*/
type Expense struct {
	gorm.Model
	Title  string `gorm:"index;not null"`
	Price  uint   `gorm:"default:0"`
	UserID uint   `gorm:"foreignkey:ID"`
}
