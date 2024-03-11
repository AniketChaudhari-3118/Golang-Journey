//Gorilla Mux

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/home", Home)
	router.HandleFunc("/user", User)
	router.HandleFunc("/do", validate(dosomething))

	router.Use(checkNumberOfRequests)
	http.ListenAndServe(":8080", router)
}

func Home(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Home function was called")
	res.Write([]byte("You called the home function"))
}

func User(res http.ResponseWriter, req *http.Request) {
	fmt.Println("New User")
	fmt.Fprintf(res, "You are the new user")
}

func dosomething(res http.ResponseWriter, req *http.Request) {
	fmt.Println("did a thing")
}

var numberOfRequests = 0

// This function is a middleware function which is used to call before all the functions
// It is mostly used for authentication of the user and etc
// Middleware functions in Go are a way to intercept and manipulate HTTP requests and responses
// before they are handled by the main request handler
// A middleware simply takes a http.HandlerFunc as one of its parameters,
// wraps it and returns a new http.HandlerFunc for the server to call.
func checkNumberOfRequests(h http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		numberOfRequests = numberOfRequests + 1
		fmt.Println("request number: ", numberOfRequests)
		h.ServeHTTP(res, req)
	})
}

// This function is also a middleware function but used to call before only a particular function.
func validate(h http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("user has been vaidated")

		h.ServeHTTP(res, req)
	}
}
