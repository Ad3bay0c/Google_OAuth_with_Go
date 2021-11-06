package controllers

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
)

func init() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/admin", Admin)
}

func Index(rw http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	u := user.Current(ctx)
	url, _ := user.LogoutURL(ctx, "/")
	rw.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(rw, "Welcome %s! <br>", u.Email)
	fmt.Fprintf(rw, "You are an Admin %v! <br>", u.Admin)

	if u.Admin {
		fmt.Fprintf(rw, "(<a href='/admin'>go toc admin</a>) <br>")
		
	}
	fmt.Fprintf(rw, "(<a href='%s'>signout</a>) <br>", url)
}

func Admin(rw http.ResponseWriter, req *http.Request) {

}
