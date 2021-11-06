package main

import (
	"log"
	"net/http"

	"github.com/Ad3bay0c/go-auth/controllers"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(err.Error())
	}

	// router := pat.New()
	// router.Get("/api/auth/{provider}", controllers.Login)
	// router.Get("/", controllers.LoginPage)
	// router.Get("/auth/{provider}/callback", controllers.OAuthCallback)

	// log.Println("Server Started on localhost:3500")

	// // client := db.RedisConnect()

	// // fmt.Println(client.Get("name"))
	// log.Fatal(http.ListenAndServe(":3500", router))

	mux := http.NewServeMux()
	// Root

	// OauthGoogle
	mux.HandleFunc("/api/auth/google", controllers.Login)
	mux.HandleFunc("/", controllers.LoginPage)
	mux.HandleFunc("/auth/google/callback", controllers.OAuthCallback)
	//
	PORT := ":3500"
	server := &http.Server{
		Handler: mux,
		Addr:    PORT,
	}
	server.ListenAndServe()
}
