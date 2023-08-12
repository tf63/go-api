// /*
// # RESTのサーバー
//   - Ginでサーバーを起動する
//   - エンドポイント
//   - /rest/todos/1 -> userIDが1のTodoをGET
//   - /rest/todos -> TodoのPOST, PUT, DELETE
//   - /rest/users/1 -> userIDが1のUserをGET
//   - /rest/users/ -> UserのPOST
//
// */
package main

import (
	"flag"
	"log"
	"net"
	"net/http"

	"github.com/tf63/go_api/api/rest"
	"github.com/tf63/go_api/external"
	handler "github.com/tf63/go_api/internal/handler/rest"
	"github.com/tf63/go_api/internal/repository"
)

// import (
// 	"github.com/gin-contrib/cors"
// 	"github.com/gin-gonic/gin"
// 	"github.com/tf63/go_api/external"
// )

func main() {

	db, _ := external.ConnectDatabase(false)
	er := repository.NewExpenseRepository(*db)
	sh := handler.NewServerHandler(er)
	options := rest.ChiServerOptions{BaseURL: "/api"}
	ro := rest.HandlerWithOptions(sh, options)

	port := flag.String("port", "8080", "Port for test HTTP server")
	flag.Parse()

	s := &http.Server{
		Handler: ro,
		Addr:    net.JoinHostPort("0.0.0.0", *port),
	}

	log.Printf(`Starting Server... at ` + *port)
	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
