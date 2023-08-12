package repository

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tf63/go_api/api/rest"
	"github.com/tf63/go_api/external"
	"github.com/tf63/go_api/internal/entity"
)

// repository
var er ExpenseRepository
var ur UserRepository

/*
テスト本体
*/
func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

/*
テスト前の準備
*/
func setUp() {
	db, _ := external.ConnectDatabase(true)
	er = NewExpenseRepository(*db)
	ur = NewUserRepository(*db)
}

func Test(t *testing.T) {
	// ----------------------------------------------------------------
	// Test OK (200)
	// ----------------------------------------------------------------

	// 前データ
	user1Name := "user1"
	user1 := rest.NewUser{Name: &user1Name}
	user2Name := "user1"
	user2 := rest.NewUser{Name: &user2Name}
	user3Name := "user3"
	user3 := rest.NewUser{Name: &user3Name}

	findUser1UserId := 1
	findUser1 := rest.FindUser{UserId: &findUser1UserId}
	findUser2UserId := 2
	findUser2 := rest.FindUser{UserId: &findUser2UserId}
	findUser3UserId := 3
	findUser3 := rest.FindUser{UserId: &findUser3UserId}

	expense1Title := "user1's expense 1"
	expense1Price := 100
	expense1UserId := 1
	expense1 := rest.NewExpense{Title: &expense1Title, Price: &expense1Price, UserId: &expense1UserId}
	expense2Title := "user1's expense 2"
	expense2Price := 200
	expense2UserId := 1
	expense2 := rest.NewExpense{Title: &expense2Title, Price: &expense2Price, UserId: &expense2UserId}
	expense3Title := "user2's expense 1"
	expense3Price := 1000
	expense3UserId := 2
	expense3 := rest.NewExpense{Title: &expense3Title, Price: &expense3Price, UserId: &expense3UserId}

	updatedExpense1Title := "updated user1's expense 1"
	updatedExpense1 := rest.NewExpense{Title: &updatedExpense1Title, UserId: &findUser1UserId}

	updatedExpense2Title := "updated user2's expense 1"
	updatedExpense2 := rest.NewExpense{Title: &updatedExpense2Title, UserId: &findUser2UserId}

	updatedExpense3Title := "updated user3's expense 1"
	updatedExpense3 := rest.NewExpense{Title: &updatedExpense3Title, UserId: &findUser3UserId}

	// Create: ユーザーを作成
	if true {

		userId, err := ur.CreateUser(user1)
		assert.Equal(t, nil, err)
		assert.Equal(t, 1, userId)
		userId, err = ur.CreateUser(user2)
		assert.Equal(t, nil, err)
		assert.Equal(t, 2, userId)
		userId, err = ur.CreateUser(user3)
		assert.Equal(t, nil, err)
		assert.Equal(t, 3, userId)
	}

	// Read: 作成できているか確認
	if true {
		result, err := ur.ReadUser(1)
		assert.Equal(t, *user1.Name, result.Name)
		assert.Equal(t, nil, err)
		result, err = ur.ReadUser(2)
		assert.Equal(t, *user2.Name, result.Name)
		assert.Equal(t, nil, err)
		result, err = ur.ReadUser(3)
		assert.Equal(t, *user3.Name, result.Name)
		assert.Equal(t, nil, err)
	}

	// Create: Expenseを作成
	if true {
		expenseId, err := er.CreateExpense(expense1)
		assert.Equal(t, nil, err)
		assert.Equal(t, 1, expenseId)
		expenseId, err = er.CreateExpense(expense2)
		assert.Equal(t, nil, err)
		assert.Equal(t, 2, expenseId)
		expenseId, err = er.CreateExpense(expense3)
		assert.Equal(t, nil, err)
		assert.Equal(t, 1, expenseId)
	}

	// Read: 作成できているか確認
	if true {
		result, err := er.ReadExpenses(findUser1)
		assert.Equal(t, 2, len(result))
		assert.Equal(t, nil, err)
		result, err = er.ReadExpenses(findUser2)
		assert.Equal(t, 1, len(result))
		assert.Equal(t, nil, err)
		result, err = er.ReadExpenses(findUser3)
		assert.Equal(t, 0, len(result))
		assert.Equal(t, nil, err)
	}

	// Update: Expenseを更新
	if true {
		err := er.UpdateExpense(updatedExpense1, 1)
		assert.Equal(t, nil, err)
		err = er.UpdateExpense(updatedExpense2, 1)
		assert.Equal(t, nil, err)
		// ユーザー1のExpenseが更新できているか確認
		result, err := er.ReadExpense(findUser1, 1)
		assert.Equal(t, *updatedExpense1.Title, result.Title) // 1番目のExpenseを更新できている?
		assert.Equal(t, nil, err)
		// ユーザー2のExpenseが更新できているか確認
		result, err = er.ReadExpense(findUser2, 1)
		assert.Equal(t, *updatedExpense2.Title, result.Title) // 1番目のExpenseを更新できている?
		assert.Equal(t, nil, err)
	}

	// Delete: Expenseを削除
	if true {
		err := er.DeleteExpense(findUser1, 1)
		assert.Equal(t, nil, err)
		err = er.DeleteExpense(findUser2, 1)
		assert.Equal(t, nil, err)
		// ユーザー1のExpenseが削除できているか確認
		result, err := er.ReadExpenses(findUser1)
		assert.Equal(t, *expense2.Title, result[0].Title) // 1番目を削除した結果，残っているのはUpdateしてないやつ?
		assert.Equal(t, nil, err)
		result, err = er.ReadExpenses(findUser2)
		assert.Equal(t, 0, len(result)) // ユーザー2のデータは残っていない?
		assert.Equal(t, nil, err)
	}

	// ----------------------------------------------------------------
	// Test Not Found (404)
	// ----------------------------------------------------------------

	// Update: 存在しないExpenseを更新すると404エラー
	if true {
		err := er.UpdateExpense(updatedExpense1, 2000)
		assert.Equal(t, entity.STATUS_NOT_FOUND, err)
		err = er.UpdateExpense(updatedExpense2, 1) // ユーザー1のExpenseID=1は存在しているが，ユーザー2は?
		assert.Equal(t, entity.STATUS_NOT_FOUND, err)
		err = er.UpdateExpense(updatedExpense3, 1) // ユーザー3は元々何もない
		assert.Equal(t, entity.STATUS_NOT_FOUND, err)
	}

	// Delete: 存在しないExpenseを削除すると404エラー
	if true {
		err := er.DeleteExpense(findUser1, 2000)
		assert.Equal(t, entity.STATUS_NOT_FOUND, err)
		err = er.DeleteExpense(findUser2, 1) // ユーザー1のExpenseID=1は存在しているが，ユーザー2は?
		assert.Equal(t, entity.STATUS_NOT_FOUND, err)
		err = er.DeleteExpense(findUser3, 1) // ユーザー3は元々何もない
		assert.Equal(t, entity.STATUS_NOT_FOUND, err)
	}

	// ----------------------------------------------------------------
	// Test Service Not Unavailable  (503) ここでは起こせないものがある
	// ----------------------------------------------------------------

	// CRUD: 適当なエラー
	// user_idがnilになるとエラーを吐く
	if true {
		_, err := er.CreateExpense(rest.NewExpense{}) // nil
		assert.Equal(t, entity.STATUS_SERVICE_UNAVAILABLE, err)
		_, err = er.ReadExpense(rest.FindUser{}, 1) // nil
		assert.Equal(t, entity.STATUS_SERVICE_UNAVAILABLE, err)
		_, err = er.ReadExpenses(rest.FindUser{})
		assert.Equal(t, entity.STATUS_SERVICE_UNAVAILABLE, err)
		err = er.UpdateExpense(rest.NewExpense{}, 1)
		assert.Equal(t, entity.STATUS_SERVICE_UNAVAILABLE, err)
		err = er.DeleteExpense(rest.FindUser{}, 1)
		assert.Equal(t, entity.STATUS_SERVICE_UNAVAILABLE, err)
	}

	// Create: 適当なエラー
	if true {
		_, err := ur.CreateUser(rest.NewUser{}) // nil
		assert.Equal(t, entity.STATUS_SERVICE_UNAVAILABLE, err)
	}
}
