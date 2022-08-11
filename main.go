package main

import (
	conn "financial-api/m/db"
	c "financial-api/m/graphql/check_api"
	m "financial-api/m/graphql/mileage_api"
	p "financial-api/m/graphql/petty_api"
	u "financial-api/m/graphql/user_api"
	"net/http"
	"os"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

const defaultPort = "8080"

var userSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    u.UserQueries,
	Mutation: u.UserMutations,
})
var mileageSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    m.MileageQueries,
	Mutation: m.MileageMutations,
})
var pettyCashSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    p.PettyCashQueries,
	Mutation: p.PettyCashMutations,
})
var checkRequestSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    c.CheckQueries,
	Mutation: c.CheckRequestMutations,
})

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	conn.InitDB()
	defer conn.CloseDB()
	userHandler := handler.New(&handler.Config{
		Schema:     &userSchema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})
	mileageHandler := handler.New(&handler.Config{
		Schema:     &mileageSchema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})
	pettyCashHandler := handler.New(&handler.Config{
		Schema:     &pettyCashSchema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})
	checkRequestHandler := handler.New(&handler.Config{
		Schema:     &checkRequestSchema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})
	http.Handle("/user", userHandler)
	http.Handle("/mileage", mileageHandler)
	http.Handle("/petty_cash", pettyCashHandler)
	http.Handle("/check_request", checkRequestHandler)
	http.ListenAndServe(":8080", nil)
}
