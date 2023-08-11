// /*
// # RESTのサーバー
//   - Ginでサーバーを起動する
//   - エンドポイント
//   - /rest/todos/1 -> userIDが1のTodoをGET
//   - /rest/todos -> TodoのPOST, PUT, DELETE
//   - /rest/users/1 -> userIDが1のUserをGET
//   - /rest/users/ -> UserのPOST

// 参考 https://github.com/koga456/sample-api/blob/master/cmd/sample-api/main.go
// */
package main

// import (
// 	"github.com/gin-contrib/cors"
// 	"github.com/gin-gonic/gin"
// 	"github.com/tf63/go_api/external"
// )

// // Ginのサーバーを起動
// func main() {
// 	// レイヤードアーキテクチャでAPIを設計
// 	// DIを入れている
// 	db, _ := external.ConnectDatabase(false)
// 	tr := repository.NewTodoRepository(*db)
// 	ur := repository.NewUserRepository(*db)
// 	tc := rest.NewTodoController(tr)
// 	uc := rest.NewUserController(ur)
// 	ro := rest.NewRouter(tc, uc)

// 	// Ginのエンジンの初期化
// 	r := gin.Default()

// 	config := cors.DefaultConfig()
// 	config.AllowOrigins = []string{"http://localhost:5173"} // 許可するオリジンを指定
// 	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
// 	config.AllowHeaders = []string{"Origin", "Content-Type"}

// 	// CORSの設定
// 	r.Use(cors.New(config))

// 	// routingの設定
// 	r.GET("/rest/todos/:user-id", ro.HandleTodosRequest)
// 	r.POST("/rest/todos", ro.HandleTodosRequest)
// 	r.PUT("/rest/todos", ro.HandleTodosRequest)
// 	r.DELETE("/rest/todos", ro.HandleTodosRequest)

// 	r.GET("/rest/users/:user-id", ro.HandleUsersRequest)
// 	r.POST("/rest/users", ro.HandleUsersRequest)

// 	// Ginのサーバーの起動
// 	err := r.Run(":8080")
// 	if err != nil {
// 		panic(err)
// 	}
// }
