package main

import (
	conn "financial-api/db"
	r "financial-api/graphql/root"
	auth "financial-api/middleware"
	"fmt"
	"net/http"
	"os"
	// "encoding/json"
	"github.com/go-chi/chi"
	"github.com/gorilla/handlers"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

const defaultPort = "8080"

var rootSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    r.RootQueries,
	Mutation: r.RootMutations,
})

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema: schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
	return result
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	conn.InitDB()
	defer conn.CloseDB()
	rootRequestHandler := handler.New(&handler.Config{
		Schema:     &rootSchema,
		Pretty:     true,
		GraphiQL:   true,
	})
	router := chi.NewRouter()
	router.Use(auth.Middleware())
	router.Handle("/graphql", rootRequestHandler)
	http.Handle("/graphql", rootRequestHandler)
	http.HandleFunc("/", rootRequestHandler)
	originsOK := handlers.AllowedOrigins([]string{"https://agile-tundra-78417.herokuapp.com/graphql", "http://localhost:3000", "https://finance-requests.vercel.app"})
	headersOK := handlers.AllowedHeaders([]string{"Content-Type", "Authorization", "X-Requested-With"})
	methodsOK := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	http.ListenAndServe(":"+port, handlers.CORS(originsOK, headersOK, methodsOK)(router))
}
