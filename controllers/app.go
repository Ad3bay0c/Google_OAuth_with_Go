package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"html/template"

	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"golang.org/x/oauth2"

	// "log"
	"net/http"
	"os"
	"time"
)

func BeginAuth(rw http.ResponseWriter, req *http.Request) {
	gothic.BeginAuthHandler(rw, req)
}

func LoginPage(rw http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("./templates/login.gohtml")
	t.Execute(rw, nil)

}

func Callback(rw http.ResponseWriter, req *http.Request) {
	user, err := gothic.CompleteUserAuth(rw, req)
	if err != nil {
		fmt.Fprintf(rw, "Error : %v", err.Error())
		return
	}

	fmt.Fprintf(rw, "Data %v", user)
	// log.Println(user)
	// t, _ := template.ParseFiles("./templates/profile.gohtml")
	// t.Execute(rw, user)
}

var googleOAuthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:3500/auth/google/callback",
	ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

func OAuthLogin(w http.ResponseWriter, req *http.Request) {
	// Create Oauth State Cookie
	oauthState := generateStateOauthCookie(w)
	u := googleOAuthConfig.AuthCodeURL(oauthState)
	http.Redirect(w, req, u, http.StatusTemporaryRedirect)
}

func generateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}
