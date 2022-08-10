package main

import (
	"encoding/json"
	c "financial-api/m/graphql/check_api"
	m "financial-api/m/graphql/mileage_api"
	p "financial-api/m/graphql/petty_api"
	u "financial-api/m/graphql/user_api"
	"fmt"
	"net/http"

	"github.com/friendsofgo/graphiql"
	"github.com/graphql-go/graphql"
)

var rootQueries = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"user": &graphql.Field{
			Name: "User Queries",
			Type: u.UserQueries,
		},
		"mileage": &graphql.Field{
			Name: "Mileage Queries",
			Type: m.MileageQueries,
		},
		"petty_cash": &graphql.Field{
			Name: "Petty Cash Queries",
			Type: p.PettyCashQueries,
		},
		"check": &graphql.Field{
			Name: "Check Request Queries",
			Type: c.CheckQueries,
		},
	},
})

var rootMutations = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"user": &graphql.Field{
			Name: "User Mutations",
			Type: u.UserMutations,
		},
		"mileage": &graphql.Field{
			Name: "Mileage Mutations",
			Type: m.MileageMutations,
		},
		"petty_cash": &graphql.Field{
			Name: "Petty Cash Mutations",
			Type: p.PettyCashMutations,
		},
		"check": &graphql.Field{
			Name: "Check Request Mutations",
			Type: c.CheckRequestMutations,
		},
	},
})

var rootSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQueries,
	Mutation: rootMutations,
})

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        rootSchema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
	return result
}

func main() {
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), rootSchema)
		json.NewEncoder(w).Encode(result)
	})
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
	// below modification allows for removal of graphiql on deployment
	graphiqlHandler, err := graphiql.NewGraphiqlHandler("http://localhost:8080/graphql")
	if err != nil {
		panic(err)
	}
	http.Handle("/graphiql", graphiqlHandler)
	http.ListenAndServe(":4040", nil)
}
