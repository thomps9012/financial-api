package main

import (
	"encoding/json"
	c "financial-api/m/checkAPI"
	m "financial-api/m/mileageAPI"
	p "financial-api/m/pettyAPI"
	u "financial-api/m/userAPI"
	"fmt"
	"net/http"

	"github.com/friendsofgo/graphiql"
	"github.com/graphql-go/graphql"
)

var rootQueries = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"user":       u.UserQueries,
		"mileage":    m.MileageQueries,
		"petty_cash": p.PettyCashQueries,
		"check":      c.CheckQueries,
	},
})

var rootMutations = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"user":       u.UserMutations,
		"mileage":    m.MileageMutations,
		"petty_cash": p.PettyCashMutations,
		"check":      c.CheckMutations,
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
