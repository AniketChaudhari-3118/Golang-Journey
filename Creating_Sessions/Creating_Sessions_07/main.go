package main

import (
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName string
	Password []byte
	First    string
	last     string
	Role     string
}

var tpl *template.Template
var dbUsers = map[string]user{}      // user ID, user
var dbSessions = map[string]string{} //session ID, user ID

func init() {
	tpl = template.Must(template.ParseGlob("template/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	u := getUser(res, req)
	tpl.ExecuteTemplate(res, "index (3).gohtml", u)
}

func bar(res http.ResponseWriter, req *http.Request) {
	u := getUser(res, req)
	if !alreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	if u.Role != "007" {

		http.Error(res, "You must be 007 to enter the bar", http.StatusForbidden)
		return
	}
	tpl.ExecuteTemplate(res, "bar.gohtml", u)

}

func signup(res http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	//process form submission
	if req.Method == http.MethodPost {

		//get form values
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		r := req.FormValue("role")
		//usernametaken?
		if _, ok := dbUsers[un]; ok {
			http.Error(res, "Username already taken", http.StatusForbidden)
			return
		}

		//create session
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(res, c)
		dbSessions[c.Value] = un

		//store users in dbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(res, "Internal Server Error: ", http.StatusInternalServerError)
			return
		}
		u := user{un, bs, f, l, r}
		dbUsers[un] = u
		//redirect
		http.Redirect(res, req, "/", http.StatusSeeOther)
	}

	tpl.ExecuteTemplate(res, "signup.gohtml", nil)
}

func login(res http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	//process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		//is their  a username?
		u, ok := dbUsers[un]
		if !ok {
			http.Error(res, "Username and/or Password do not match", http.StatusForbidden)
			return
		}

		//does the entered password match the stored password
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(res, "Username and/or Password do not match", http.StatusForbidden)
			return
		}

		//create session
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(res, c)
		dbSessions[c.Value] = un
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(res, "login.gohtml", nil)
}

func logout(res http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	c, _ := req.Cookie("session")
	//delete the session
	delete(dbSessions, c.Value)
	//remove the cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(res, c)

	http.Redirect(res, req, "/login", http.StatusSeeOther)
}
