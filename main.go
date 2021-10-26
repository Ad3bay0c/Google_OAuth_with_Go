package main

import (
	"github.com/Ad3bay0c/go-auth/controllers"
	"github.com/gorilla/pat"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(err.Error())
	}

	maxAge := 86400 * 30  // 30 days
	isProd := false       // Set to true when serving over https

	store := sessions.NewCookieStore([]byte(os.Getenv("SECRET_KEY")))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true   // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.Store = store

	goth.UseProviders(
		google.New(os.Getenv("SECRET_ID"), os.Getenv("SECRET_KEY"),
			"http://localhost:3500/auth/google/callback"),
	)

	router := pat.New()
	router.Get("/auth/{provider}", controllers.BeginAuth)
	router.Get("/", controllers.LoginPage)
	router.Get("/auth/{provider}/callback", controllers.Callback)

	log.Println("Server Started on localhost:3500")
	log.Fatal(http.ListenAndServe(":3500", router))
}