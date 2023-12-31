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

func main() {

	db, err := external.ConnectDatabase()
	if err != nil {
		log.Fatal("Failed to Connect Database")
	}

	er := repository.NewExpenseRepository(*db)
	ur := repository.NewUserRepository(*db)
	sh := handler.NewServerHandler(er, ur)
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
