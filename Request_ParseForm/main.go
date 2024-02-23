package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type hotdog int

// func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
// 	err := req.ParseForm()
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	err = tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method        string
		URL           *url.URL
		Submissions   map[string][]string
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
		req.Host,
		req.ContentLength,
	}

	err = tpl.ExecuteTemplate(w, "index.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}

}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	log.Fatal(http.ListenAndServe(":8082", d))
}
