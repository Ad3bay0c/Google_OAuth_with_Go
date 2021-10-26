package controllers

import (
	"fmt"
	"github.com/markbates/goth/gothic"
	"html/template"
	"log"
	"net/http"
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
	log.Println(user)
	t, _ := template.ParseFiles("./templates/profile.gohtml")
	t.Execute(rw, user)
}