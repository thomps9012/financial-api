package main

import (
	conn "financial-api/db"
	c "financial-api/graphql/check_api"
	m "financial-api/graphql/mileage_api"
	p "financial-api/graphql/petty_api"
	u "financial-api/graphql/user_api"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/pat"
	"github.com/gorilla/sessions"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
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

// var (

// 	oauthConfig = &oauth2.Config{
// 		ClientID:     os.Getenv("GOOGLE_OAUTH_ID"),
// 		ClientSecret: os.Getenv("GOOGLE_OAUTH_SECRET"),
// 		Endpoint:     google.Endpoint,
// 		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
// 		RedirectURL:  "https://" + os.Getenv("HEROKU_APP_NAME") + "herokuapp.com/auth/google/callback",
// 	}
// 	oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
// )

// func handleRoot(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, `<html><body><a href="/auth/google">Sign In with Google</a></body></html>`)
// }

// func handleAuth(w http.ResponseWriter, r *http.Request) {
// 	oauthState := generateOauthCookie(w)
// 	u := oauthConfig.AuthCodeURL(oauthState)
// 	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
// }

// func generateOauthCookie(w http.ResponseWriter) string {
// 	var expiration = time.Now().Add(20 * time.Minute)

// 	b := make([]byte, 16)
// 	rand.Read(b)
// 	state := base64.URLEncoding.EncodeToString(b)
// 	cookie := http.Cookie{Name: "tundra-oauth", Value: state, Expires: expiration}
// 	http.SetCookie(w, &cookie)

// 	return state
// }

// func handleAuthCallback(w http.ResponseWriter, r *http.Request) {
// 	oauthState, _ := r.Cookie("tundra-oauth")
// 	if r.FormValue("state") != oauthState.Value {
// 		log.Println("invalid oauth google state")
// 		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
// 		return
// 	}

// 	data, err := handleUser(r.FormValue("code"))
// 	if err != nil {
// 		log.Println(err.Error())
// 		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
// 		return
// 	}
// 	fmt.Fprintf(w, `<html><body><h1>Hello %s\n</h1><br /><h2>Below are some of the available APIs for the application</h2><br/><a href="/user_api">User Management API</a><br/><a href="/mileage_api">Mileage Request API</a><br/><a href="/petty_cash_api">Petty Cash API</a><br/><a href="/check_request_api">Check Request API</a></body></html>`, data)
// }

// func handleUser(code string) ([]byte, error) {
// 	// Use code to get token and get user info from Google.

// 	token, err := oauthConfig.Exchange(context.Background(), code)
// 	if err != nil {
// 		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
// 	}
// 	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
// 	}
// 	defer response.Body.Close()
// 	contents, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed read response: %s", err.Error())
// 	}
// 	return contents, nil
// }
func main() {
	maxAge := 86400 * 1
	isProd := true
	store := sessions.NewCookieStore([]byte(os.Getenv("COOKIE_SECRET")), []byte(os.Getenv("COOKIE_ENCRYPT")))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = isProd

	gothic.Store = store
	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_OAUTH_ID"), os.Getenv("GOOGLE_OAUTH_SECRET"), "https://"+os.Getenv("HEROKU_APP_NAME")+"herokuapp.com/auth/google/callback", "email", "profile"),
	)
	path := pat.New()
	path.Get("/auth/${provider}/callback", func(res http.ResponseWriter, req *http.Request) {

		user, err := gothic.CompleteUserAuth(res, req)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.ParseFiles("templates/success.html")
		t.Execute(res, user)
		// fmt.Fprintf(`<html><body><h1>Hello %s\n</h1><br /><h2>Below are some of the available APIs for the application</h2><br/><a href="/user_api">User Management API</a><br/><a href="/mileage_api">Mileage Request API</a><br/><a href="/petty_cash_api">Petty Cash API</a><br/><a href="/check_request_api">Check Request API</a></body></html>`, user)
	})
	path.Get("/auth/{provider}", func(res http.ResponseWriter, req *http.Request) {
		gothic.BeginAuthHandler(res, req)
	})

	path.Get("/", func(res http.ResponseWriter, req *http.Request) {
		t, _ := template.ParseFiles("templates/index.html")
		t.Execute(res, false)
	})
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
	// http.HandleFunc("/", handleRoot)
	// http.HandleFunc("/auth/google", handleAuth)
	// http.HandleFunc("/auth/google/callback", handleAuthCallback)
	http.Handle("/user_api", userHandler)
	http.Handle("/mileage_api", mileageHandler)
	http.Handle("/petty_cash_api", pettyCashHandler)
	http.Handle("/check_request_api", checkRequestHandler)
	http.ListenAndServe(":"+port, nil)
}
