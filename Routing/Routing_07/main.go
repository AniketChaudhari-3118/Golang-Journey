package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"
)

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Dog is barking")
}

func blank(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Nothing here!! ")
}

func name(res http.ResponseWriter, req *http.Request) {
	// io.WriteString(res, "Aniket here!")
	tpl, err := template.ParseFiles("something.gohtml")
	if err != nil {
		log.Fatalln("2001, File not Prased succesfully")
	}
	err = tpl.ExecuteTemplate(res, "something.gohtml", "Aniket")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	http.Handle("/", http.HandlerFunc(blank))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/name", http.HandlerFunc(name))

	http.ListenAndServe(":8080", nil)
}

func HandlerFunc(blank func(res http.ResponseWriter, req *http.Request)) {
	panic("unimplemented")
}
