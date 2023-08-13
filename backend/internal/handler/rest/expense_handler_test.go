package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tf63/go_api/api/rest"
	"github.com/tf63/go_api/internal/entity"
	"github.com/tf63/go_api/tests"
)

func getTestNewExpense() rest.NewExpense {
	title := "Test Expense"
	price := 100
	userId := 1

	return rest.NewExpense{
		Title:  &title,
		Price:  &price,
		UserId: &userId,
	}
}

func getTestExpenseEntity() entity.Expense {
	return entity.Expense{
		Title:  "Test Expense",
		Price:  100,
		UserID: 1,
	}
}

func getTestExpensesEntity() []entity.Expense {
	expenses := []entity.Expense{}

	expenses = append(expenses, entity.Expense{Title: "Test Expense 1", Price: 100, UserID: 1})
	expenses = append(expenses, entity.Expense{Title: "Test Expense 2", Price: 200, UserID: 2})
	expenses = append(expenses, entity.Expense{Title: "Test Expense 3", Price: 300, UserID: 3})
	return expenses
}

func getTestFindUser() rest.FindUser {
	userId := 1

	return rest.FindUser{
		UserId: &userId,
	}
}

func TestPostV1Expenses(t *testing.T) {
	// モックを作成
	mer := new(tests.MockExpenseRepository)

	// モックに期待するメソッド呼び出しを設定
	body := getTestNewExpense()
	expectedExpenseId := 123
	input, err := NewExpenseDTO(&body)
	assert.NoError(t, err)
	mer.On("CreateExpense", input).Return(expectedExpenseId, nil)

	// テストリクエストを作成
	bodyJSON, _ := json.Marshal(body)
	req, err := http.NewRequest("POST", "/v1/expenses", bytes.NewBuffer(bodyJSON))
	assert.NoError(t, err)

	// テストレスポンスを作成
	rr := httptest.NewRecorder()

	// テスト対象のハンドラを呼び出し
	sh := &serverHandler{er: mer}
	http.HandlerFunc(sh.PostV1Expenses).ServeHTTP(rr, req)

	// アサーション
	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Contains(t, rr.Header().Get("Location"), strconv.Itoa(expectedExpenseId))

	// モックのアサーション
	mer.AssertExpectations(t)
}

func TestGetV1Expenses(t *testing.T) {
	// モックを作成
	mer := new(tests.MockExpenseRepository)

	// モックに期待するメソッド呼び出しを設定
	body := getTestFindUser()
	expectedExpenses := getTestExpensesEntity()
	mer.On("ReadExpenses", FindUserDTO(&body)).Return(expectedExpenses, nil)

	// テストリクエストを作成
	bodyJSON, _ := json.Marshal(body)
	req, err := http.NewRequest("GET", "/v1/expenses", bytes.NewBuffer(bodyJSON))
	assert.NoError(t, err)

	// テストレスポンスを作成
	rr := httptest.NewRecorder()

	// テスト対象のハンドラを呼び出し
	sh := &serverHandler{er: mer}
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sh.GetV1Expenses(w, r)
	}).ServeHTTP(rr, req)

	// アサーション
	assert.Equal(t, http.StatusOK, rr.Code)

	// レスポンスボディをパースして検証
	var response rest.Expenses
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Len(t, response, len(expectedExpenses))

	// モックのアサーション
	mer.AssertExpectations(t)
}

func TestDeleteV1ExpensesExpenseId(t *testing.T) {
	// モックを作成
	mer := new(tests.MockExpenseRepository)

	// モックに期待するメソッド呼び出しを設定
	body := getTestFindUser()
	expectedExpenseId := 123
	mer.On("DeleteExpense", FindUserDTO(&body), expectedExpenseId).Return(nil)

	// テストリクエストを作成
	bodyJSON, _ := json.Marshal(body)
	req, err := http.NewRequest("DELETE", "/v1/expenses/"+strconv.Itoa(expectedExpenseId), bytes.NewBuffer(bodyJSON))
	assert.NoError(t, err)

	// テストレスポンスを作成
	rr := httptest.NewRecorder()

	// テスト対象のハンドラを呼び出し
	sh := &serverHandler{er: mer}
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sh.DeleteV1ExpensesExpenseId(w, r, expectedExpenseId)
	}).ServeHTTP(rr, req)

	// アサーション
	assert.Equal(t, http.StatusNoContent, rr.Code)

	// モックのアサーション
	mer.AssertExpectations(t)
}

func TestGetV1ExpensesExpenseId(t *testing.T) {
	// モックを作成
	mer := new(tests.MockExpenseRepository)

	// モックに期待するメソッド呼び出しを設定
	body := getTestFindUser()
	expectedExpenseId := 123
	expectedExpense := getTestExpenseEntity()

	mer.On("ReadExpense", FindUserDTO(&body), expectedExpenseId).Return(expectedExpense, nil)

	// テストリクエストを作成
	bodyJSON, _ := json.Marshal(body)
	req, err := http.NewRequest("GET", "/v1/expenses/"+strconv.Itoa(expectedExpenseId), bytes.NewBuffer(bodyJSON))
	assert.NoError(t, err)

	// テストレスポンスを作成
	rr := httptest.NewRecorder()

	// テスト対象のハンドラを呼び出し
	sh := &serverHandler{er: mer}
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sh.GetV1ExpensesExpenseId(w, r, expectedExpenseId)
	}).ServeHTTP(rr, req)

	// アサーション
	assert.Equal(t, http.StatusOK, rr.Code)

	// レスポンスボディをパースして検証
	var response rest.Expense
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, expectedExpense.Title, *response.Title)
	assert.Equal(t, expectedExpense.Price, uint(*response.Price))
	assert.Equal(t, expectedExpense.UserID, uint(*response.UserId))

	// モックのアサーション
	mer.AssertExpectations(t)
}

func TestPutV1ExpensesExpenseId(t *testing.T) {
	// モックを作成
	mer := new(tests.MockExpenseRepository)

	// モックに期待するメソッド呼び出しを設定
	body := getTestNewExpense()
	expectedExpenseId := 123
	input, err := NewExpenseDTO(&body)
	assert.NoError(t, err)
	mer.On("UpdateExpense", input, expectedExpenseId).Return(nil)

	// テストリクエストを作成
	bodyJSON, _ := json.Marshal(body)
	req, err := http.NewRequest("PUT", "/v1/expenses/"+strconv.Itoa(expectedExpenseId), bytes.NewBuffer(bodyJSON))
	assert.NoError(t, err)

	// テストレスポンスを作成
	rr := httptest.NewRecorder()

	// テスト対象のハンドラを呼び出し
	sh := &serverHandler{er: mer}
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sh.PutV1ExpensesExpenseId(w, r, expectedExpenseId)
	}).ServeHTTP(rr, req)

	// アサーション
	assert.Equal(t, http.StatusNoContent, rr.Code)

	// モックのアサーション
	mer.AssertExpectations(t)
}
