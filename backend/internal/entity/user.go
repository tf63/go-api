/*
# UserのEntity層
  - GORMでマイグレーションを行う
*/

package entity

import (
	"gorm.io/gorm"
)

/*
UserEntity
  - Name: ユーザー名
*/
type User struct {
	gorm.Model
	Name string `gorm:"not null"`
}
