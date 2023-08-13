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

func getTestNewUser() rest.NewUser {
	name := "Test user"

	return rest.NewUser{
		Name: &name,
	}
}

func getTestUserEntity() entity.User {
	return entity.User{
		Name: "Test User",
	}
}

func getTestUsersEntity() []entity.User {
	users := []entity.User{}

	users = append(users, entity.User{Name: "Test User 1"})
	users = append(users, entity.User{Name: "Test User 2"})
	users = append(users, entity.User{Name: "Test User 3"})
	return users
}

func TestPostV1Users(t *testing.T) {
	// モックを作成
	mur := new(tests.MockUserRepository)

	// モックに期待するメソッド呼び出しを設定
	body := getTestNewUser()
	expectedUserId := 123
	mur.On("CreateUser", NewUserDTO(&body)).Return(expectedUserId, nil)

	// テストリクエストを作成
	bodyJSON, _ := json.Marshal(body)
	req, err := http.NewRequest("POST", "/v1/users", bytes.NewBuffer(bodyJSON))
	assert.NoError(t, err)

	// テストレスポンスを作成
	rr := httptest.NewRecorder()

	// テスト対象のハンドラを呼び出し
	sh := &serverHandler{ur: mur}
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sh.PostV1Users(w, r)
	}).ServeHTTP(rr, req)

	// アサーション
	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, "/v1/users/"+strconv.Itoa(expectedUserId), rr.Header().Get("Location"))

	// モックのアサーション
	mur.AssertExpectations(t)
}

func TestGetV1Users(t *testing.T) {
	// モックを作成
	mur := new(tests.MockUserRepository)

	// モックに期待するメソッド呼び出しを設定
	expectedUsers := getTestUsersEntity()
	mur.On("ReadUsers").Return(expectedUsers, nil)

	// テストリクエストを作成
	req, err := http.NewRequest("GET", "/v1/users", nil)
	assert.NoError(t, err)

	// テストレスポンスを作成
	rr := httptest.NewRecorder()

	// テスト対象のハンドラを呼び出し
	sh := &serverHandler{ur: mur}
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sh.GetV1Users(w, r)
	}).ServeHTTP(rr, req)

	// アサーション
	assert.Equal(t, http.StatusOK, rr.Code)

	// レスポンスボディをパースして検証
	var response rest.Users
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Len(t, response, len(expectedUsers))

	// モックのアサーション
	mur.AssertExpectations(t)
}

func TestDeleteV1UsersUserId(t *testing.T) {
	// モックを作成
	mur := new(tests.MockUserRepository)

	// モックに期待するメソッド呼び出しを設定
	expectedUserId := 1
	mur.On("DeleteUser", expectedUserId).Return(nil)

	// テストリクエストを作成
	req, err := http.NewRequest("DELETE", "/v1/users/"+strconv.Itoa(expectedUserId), nil)
	assert.NoError(t, err)

	// テストレスポンスを作成
	rr := httptest.NewRecorder()

	// テスト対象のハンドラを呼び出し
	sh := &serverHandler{ur: mur}
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sh.DeleteV1UsersUserId(w, r, expectedUserId)
	}).ServeHTTP(rr, req)

	// アサーション
	assert.Equal(t, http.StatusNoContent, rr.Code)

	// モックのアサーション
	mur.AssertExpectations(t)
}

func TestGetV1UsersUserId(t *testing.T) {
	// モックを作成
	mur := new(tests.MockUserRepository)

	// モックに期待するメソッド呼び出しを設定
	expectedUserId := 1
	expectedUser := getTestUserEntity()
	mur.On("ReadUser", expectedUserId).Return(expectedUser, nil)

	// テストリクエストを作成
	req, err := http.NewRequest("GET", "/v1/users/"+strconv.Itoa(expectedUserId), nil)
	assert.NoError(t, err)

	// テストレスポンスを作成
	rr := httptest.NewRecorder()

	// テスト対象のハンドラを呼び出し
	sh := &serverHandler{ur: mur}
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sh.GetV1UsersUserId(w, r, expectedUserId)
	}).ServeHTTP(rr, req)

	// アサーション
	assert.Equal(t, http.StatusOK, rr.Code)

	// レスポンスボディをパースして検証
	var response rest.User
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser.Name, *response.Name)

	// モックのアサーション
	mur.AssertExpectations(t)
}

func TestPutV1UsersUserId(t *testing.T) {
	// モックを作成
	mur := new(tests.MockUserRepository)

	// モックに期待するメソッド呼び出しを設定
	body := getTestNewUser()
	expectedUserId := 1
	mur.On("UpdateUser", NewUserDTO(&body), expectedUserId).Return(nil)

	// テストリクエストを作成
	bodyJSON, _ := json.Marshal(body)
	req, err := http.NewRequest("PUT", "/v1/users/"+strconv.Itoa(expectedUserId), bytes.NewBuffer(bodyJSON))
	assert.NoError(t, err)

	// テストレスポンスを作成
	rr := httptest.NewRecorder()

	// テスト対象のハンドラを呼び出し
	sh := &serverHandler{ur: mur}
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sh.PutV1UsersUserId(w, r, expectedUserId)
	}).ServeHTTP(rr, req)

	// アサーション
	assert.Equal(t, http.StatusNoContent, rr.Code)

	// モックのアサーション
	mur.AssertExpectations(t)
}
