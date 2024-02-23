package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

// This piece of code present in init() will get executed as soon as all packages get imported
func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	f := req.FormValue("first")
	l := req.FormValue("last")
	s := req.FormValue("subscribed") == "on"

	err := tpl.ExecuteTemplate(res, "index.gohtml", person{f, l, s})
	if err != nil {

		http.Error(res, err.Error(), 500)
		log.Fatalln(err)
	}
}
