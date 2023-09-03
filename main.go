package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/sutin1234/go-graphql/schemas"
)

func main() {
	fmt.Println("Go GraphQL Server")

	h := handler.New(&handler.Config{
		Schema:   &schemas.BeastSchema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalf("server not start %v", err)
	}
}
