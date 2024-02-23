package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dog.jpg/", chien)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "foo ran")
}

func dog(res http.ResponseWriter, req *http.Request) {
	// res.Header().Set("Content-=Type", "Text/html; charset=utf-8")

	tpl, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		http.Error(res, "File not got Parsed", 404)
	}

	err = tpl.ExecuteTemplate(res, "dog.gohtml", "aniket")
	if err != nil {
		log.Fatal()
	}

}

func chien(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "pexels-lucas-andrade-4681107.jpg")
}

//Handson exercise
// ListenAndServe on port 8080 of localhost
// For the default route "/" Have a func called "foo" which writes to the response "foo ran"

// For the route "/dog/" Have a func called "dog" which parses a template called "dog.gohtml" and writes to the response "

// This is from dog
// " and also shows a picture of a dog when the template is executed.
// Use "http.ServeFile" to serve the file "dog.jpeg"