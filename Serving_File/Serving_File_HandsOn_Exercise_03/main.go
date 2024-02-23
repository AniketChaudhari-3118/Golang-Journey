package main

import (
	"log"
	"net/http"
)

func main() {
	fp := http.FileServer(http.Dir("public"))
	http.Handle("/", fp)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func show(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")

}
