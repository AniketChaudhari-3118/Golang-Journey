package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	fmt.Print("your request method at foo: ", req.Method, "\n\n")
}

func bar(res http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at bar is:", req.Method)
	//process from data
	// res.Header().Set("Location", "/")
	// res.WriteHeader(http.StatusSeeOther)
	http.Redirect(res, req, "/", http.StatusSeeOther)
}

func barred(res http.ResponseWriter, req *http.Request) {
	fmt.Print("your request method at barred: ", req.Method)
	tpl.ExecuteTemplate(res, "index(2).gohtml", nil)
}
