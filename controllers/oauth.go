package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Ad3bay0c/go-auth/db"
	"github.com/google/uuid"
	"github.com/markbates/goth/providers/google"
	"golang.org/x/oauth2"
)

var (
	oauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:3500/auth/google/callback",
		ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
		// Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Scopes:   []string{"https://www.googleapis.com/auth/contacts.readonly"},
		Endpoint: google.Endpoint,
	}

	client = db.RedisConnect()
)

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func Login(res http.ResponseWriter, req *http.Request) {
	oauthState := GenerateStateOauthCookie(res)
	oauthConfig.ClientID = os.Getenv("GOOGLE_OAUTH_CLIENT_ID")
	oauthConfig.ClientSecret = os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET")
	url := oauthConfig.AuthCodeURL(oauthState)
	http.Redirect(res, req, url, http.StatusTemporaryRedirect)
}

func OAuthCallback(res http.ResponseWriter, req *http.Request) {
	// oauthState, _ := req.Cookie("oauthstate")

	oauthState := client.Get("oauthstate")

	if req.FormValue("state") != oauthState.Val() {
		log.Printf("Invalid Oauth state")
		// fmt.Fprintf(res, "Error: %v", oauthState.Val())
		http.Redirect(res, req, "/", http.StatusTemporaryRedirect)
		return
	}

	token, err := oauthConfig.Exchange(context.Background(), req.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		http.Redirect(res, req, "/", http.StatusTemporaryRedirect)
		return
	}
	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		log.Println(err.Error())
		http.Redirect(res, req, "/", http.StatusTemporaryRedirect)
		return
	}
	defer response.Body.Close()
	user := struct {
		Email      string `json:"email"`
		ID         string `json:"id"`
		ProfilePic string `json:"profile_pic"`
	}{}

	content, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Println(err.Error())
		http.Redirect(res, req, "/", http.StatusTemporaryRedirect)
		return
	}
	err = json.Unmarshal(content, &user)
	if err != nil {
		log.Println(err.Error())
		http.Redirect(res, req, "/", http.StatusTemporaryRedirect)
		return
	}
	log.Println(user)

	fmt.Fprintf(res, "USerInfo %s\n", content)
	// t, _ := template.ParseFiles("./templates/profile.gohtml")
	// t.Execute(res, user)
}

func GenerateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour)

	state := uuid.New().String()
	cookie := &http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, cookie)
	client.Set("oauthstate", state, 24*time.Hour)
	return state
}
