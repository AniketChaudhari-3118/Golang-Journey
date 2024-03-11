//Gorilla Mux

package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template/*"))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler).Methods("GET") //schemes("http") || schemes("https")
	bookrouter := r.PathPrefix("/").Subrouter()
	bookrouter.HandleFunc("/hii", HomeHandler)

	r.HandleFunc("/update/{title}", updateProduct).Methods("GET") //.Methods("PUT")
	r.HandleFunc("/delete/{title}", deleteProduct).Methods("GET") //.Methods("Delete")
	r.HandleFunc("/products/{id}/{title}", productHandler).Host("www.mybookstore.com")
	// Host("www.mybookstore.com"): This part sets a constraint on the host of the incoming request.
	// It specifies that the route should only match requests with the host header www.mybookstore.com.
	// If a request is made to /books/{title} with a different host,
	// Gorilla Mux will not invoke the BookHandler for that request.
	http.ListenAndServe(":8080", r)
}

func productHandler(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	// productID := vars["id"]
	nameOfBook := vars["title"]
	//It takes parameters from url and stores in the productId variable

	tpl.ExecuteTemplate(res, "index.html", nameOfBook)
	// 	fmt.Fprintf(res, "Product ID: %s ", productID)
}

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" || req.Method == "POST" {
		fmt.Fprintf(res, "Hello")
		fmt.Fprintf(res, "Welcome to Home Page")
	}
}

func updateProduct(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	productName := vars["title"]
	fmt.Fprintf(res, "Updated the book %s Successfully", productName)
}

func deleteProduct(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	productName := vars["title"]
	fmt.Fprintf(res, "Deleted the book %s Successfully", productName)
}
