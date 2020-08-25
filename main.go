package main

import (
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/roelofjan-elsinga/graphql-rest-abstraction/queries"
)

func main() {

	// Schema represents the entry point to the GraphQL server
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"users": queries.UsersField,
				"user":  queries.UserField,
			},
		}),
	})

	if err != nil {
		panic(err)
	}

	// create a graphl-go HTTP handler with our previously defined schema
	// and we also set it to return pretty JSON output
	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	// serve a GraphQL endpoint at `/graphql`
	http.Handle("/graphql", h)

	fmt.Println("GraphQL server listening at: http://localhost:4000/graphql")

	// and serve!
	http.ListenAndServe(":4000", nil)
}
