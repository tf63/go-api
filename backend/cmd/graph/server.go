package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/tf63/go_api/api/graph"
	"github.com/tf63/go_api/external"
	resolver "github.com/tf63/go_api/internal/handler/graph"
	"github.com/tf63/go_api/internal/repository"
)

const defaultPort = "9090"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, _ := external.ConnectDatabase()
	ner := repository.NewExpenseRepository(*db)
	nur := repository.NewUserRepository(*db)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{Er: ner, Ur: nur}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
