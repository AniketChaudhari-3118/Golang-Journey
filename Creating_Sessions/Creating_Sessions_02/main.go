package main

import (
	"html/template"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}      // user ID, user
var dbSessions = map[string]string{} // session Id, user ID

func init() {
	tpl = template.Must(template.ParseGlob("template/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.Handle("favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {

	//get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(res, c)
	}

	//if the user exists already, get the user
	var u user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}

	//process form submission
	if req.Method == http.MethodGet {
		un := req.FormValue("username")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		u = user{un, f, l}
		dbSessions[c.Value] = un
		dbUsers[un] = u
	}

	err = tpl.ExecuteTemplate(res, "index (3).gohtml", u)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func bar(res http.ResponseWriter, req *http.Request) {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(res, req, "/", http.StatusSeeOther)
	}

	un, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(res, req, "/", http.StatusSeeOther)
	}
	u := dbUsers[un]
	tpl.ExecuteTemplate(res, "bar.gohtml", u)
}
