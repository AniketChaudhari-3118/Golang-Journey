package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/bar", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
func set(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "my-cookie",
		Value: "some value",
	})
	fmt.Fprintln(res, "Cookie Written, Check your browser")
	fmt.Fprintln(res, "In chrome go to dev tools/ application / cookies ")
}

func read(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {

		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprintln(res, "Your-Cookie: ", c)
}

//write and reading 1 cookie
