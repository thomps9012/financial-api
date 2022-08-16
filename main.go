package main

import (
	conn "financial-api/db"
	c "financial-api/graphql/check_api"
	m "financial-api/graphql/mileage_api"
	p "financial-api/graphql/petty_api"
	u "financial-api/graphql/user_api"
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
	// ctx := context.Background()

	// redirectURL := os.Getenv("OAUTH_CALLBACK")
	// if redirectURL == "" {
	// 	redirectURL = "https://" + os.Getenv("HEROKU_APP_NAME") + "herokuapp.com/auth/google/callback"
	// }
	// config := &oauth2.Config{
	// 	ClientID:     os.Getenv("GOOGLE_OAUTH_ID"),
	// 	ClientSecret: os.Getenv("GOOGLE_OAUTH_SECRET"),
	// 	Endpoint:     google.Endpoint,
	// 	Scopes:       []string{"email", "profile"},
	// 	RedirectURL:  redirectURL,
	// }
	// authURL := config.AuthCodeURL("state")
	// fmt.Printf("Follow the link to obtain an auth code: %s", authURL)
	// token, err := config.Exchange(ctx, "authorization-code")
	// if err != nil {
	// 	log.Fatal("config.Exchange: %v", err)
	// }
	// client := config.Client(oauth2.NoContext, token)
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	conn.InitDB()
	defer conn.CloseDB()
	userHandler := handler.New(&handler.Config{
		Schema:     &userSchema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: false,
	})
	mileageHandler := handler.New(&handler.Config{
		Schema:     &mileageSchema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: false,
	})
	pettyCashHandler := handler.New(&handler.Config{
		Schema:     &pettyCashSchema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: false,
	})
	checkRequestHandler := handler.New(&handler.Config{
		Schema:     &checkRequestSchema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: false,
	})
	http.Handle("/user_api", userHandler)
	http.Handle("/mileage_api", mileageHandler)
	http.Handle("/petty_cash_api", pettyCashHandler)
	http.Handle("/check_request_api", checkRequestHandler)
	http.ListenAndServe(":"+port, nil)
}
