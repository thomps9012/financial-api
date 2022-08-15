package main

import (
	"context"
	"encoding/gob"
	"encoding/json"
	conn "financial-api/db"
	c "financial-api/graphql/check_api"
	m "financial-api/graphql/mileage_api"
	p "financial-api/graphql/petty_api"
	u "financial-api/graphql/user_api"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
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

var (
	store = sessions.NewCookieStore([]byte(os.Getenv("COOKIE_SECRET")), []byte(os.Getenv("COOKIE_ENCRYPT")))

	oauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_OAUTH_ID"),
		ClientSecret: os.Getenv("GOOGLE_OAUTH_SECRET"),
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://" + os.Getenv("HEROKU_APP_NAME") + "herokuapp.com/.auth/login/google/callback",
	}

	stateToken = os.Getenv("HEROKU_APP_NAME")
)

func init() {
	gob.Register(&oauth2.Token{})

	store.MaxAge(60 * 60 * 8)
	store.Options.Secure = true
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<html><body><a href="/auth/google">Sign In with Google</a></body></html>`)
}

func handleAuth(w http.ResponseWriter, r *http.Request) {
	url := oauthConfig.AuthCodeURL(stateToken)
	http.Redirect(w, r, url, http.StatusFound)
}

func handleAuthCallback(w http.ResponseWriter, r *http.Request) {
	if v := r.FormValue("state"); v != stateToken {
		http.Error(w, "Invalid State token", http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	token, err := oauthConfig.Exchange(ctx, r.FormValue("code"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session, err := store.Get(r, "tundra-oauth")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["tundra-oauth"] = token
	if err := session.Save(r, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/user", http.StatusFound)
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "tundra-oauth")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	token, ok := session.Values["tundra-oauth"].(*oauth2.Token)
	if !ok {
		http.Error(w, "Unable to assert token", http.StatusInternalServerError)
		return
	}
	client := oauthConfig.Client(context.Background(), token)
	resp, err := client.Get("./auth/userinfo.email")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	d := json.NewDecoder(resp.Body)
	var account struct {
		Email string `json:"email"`
	}
	if err := d.Decode(&account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, `<html><body><h1>Hello %s</h1><br /><h2>Below are some of the available APIs for the application</h2><br/><a href="/user_api">User Management API</a><br/><a href="/mileage_api">Mileage Request API</a><br/><a href="/petty_cash_api">Petty Cash API</a><br/><a href="/check_request_api">Check Request API</a></body></html>`, account.Email)
}

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
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/auth/google", handleAuth)
	http.HandleFunc("/auth/google/callback", handleAuthCallback)
	http.HandleFunc("/user", handleUser)
	http.Handle("/user_api", userHandler)
	http.Handle("/mileage_api", mileageHandler)
	http.Handle("/petty_cash_api", pettyCashHandler)
	http.Handle("/check_request_api", checkRequestHandler)
	http.ListenAndServe(":"+port, nil)
}
