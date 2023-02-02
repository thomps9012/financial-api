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
	os.Setenv("ATLAS_URI", "mongodb+srv://spars01:H0YXCAGHoUihHcSZ@cluster0.wuezj.mongodb.net/test_finance_requests?retryWrites=true&w=majority")
	os.Setenv("DB_NAME", "test_finance_requests")
	os.Setenv("SMTP_EMAIL", "app_support@norainc.org")
	os.Setenv("SMTP_PASSWORD", "OZ8@LzR?qJU1L+Z2KNwK")
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	conn.InitDB()
	defer conn.CloseDB()
	// delete below on final production push
	rootRequestHandler := handler.New(&handler.Config{
		Schema: &rootSchema,
		Pretty: true,
	})
	router := chi.NewRouter()
	router.Use(auth.Middleware())
	router.Handle("/graphql", rootRequestHandler)
	// originsOK := handlers.AllowedOrigins([]string{"https://finance-requests.vercel.app", "http://localhost:3000"})
	originsOK := handlers.AllowedOrigins([]string{"http://localhost:3000", "http://localhost:8080/graphql", "http://localhost:8080"})
	headersOK := handlers.AllowedHeaders([]string{"Content-Type", "Authorization", "X-Requested-With"})
	methodsOK := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	http.ListenAndServe(":"+port, handlers.CORS(originsOK, headersOK, methodsOK)(auth.LimitApiCalls(router)))
}
