package main

import (
	conn "financial-api/db"
	r "financial-api/graphql/root"
	auth "financial-api/middleware"
	"net/http"
	"os"

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
		GraphiQL:   false,
		Playground: false,
	})
	router := chi.NewRouter()
	router.Use(auth.Middleware())
	router.Handle("/graphql", rootRequestHandler)
	originsOK := handlers.AllowedOrigins([]string{"*"})
	headersOK := handlers.AllowedHeaders([]string{"Content-Type", "Authorization", "X-Requested-With"})
	methodsOK := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	http.ListenAndServe(":"+port, handlers.CORS(originsOK, headersOK, methodsOK)(router))
}
