package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)


func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Get("/user-info", userInfoHandler)
	http.ListenAndServe(":3000", r)
}

func homeHandler(res http.ResponseWriter, req *http.Request) {
	ctx := make(map[string]string)
	ctx["name"] = "Aniket"

	tpl, _ := template.ParseFiles("templates/index.html")
	err := tpl.Execute(res, ctx)

	if err != nil {
		log.Println("Error in template Execution", err.Error())
	}
}

func userInfoHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("User info from Api Server"))
}
