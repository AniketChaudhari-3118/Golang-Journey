package main

import (
	"fmt"
	"log"
	"net/http"
)

var count uint64

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/bar", read)
	http.HandleFunc("/abundance", abundance)
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
	count++
}

func read(res http.ResponseWriter, req *http.Request) {
	c1, err := req.Cookie("my-cookie")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(res, "YOUR-COOKIE #1: ", c1)
	}

	c2, err := req.Cookie("general")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res, "YOUR-COOKIE #2: ", c2)
	}

	c3, err := req.Cookie("specific")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res, "YOUR-COOKIE #3: ", c3)
	}
	fmt.Println(count)
}

func abundance(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "general",
		Value: "some other value about general thinga",
	})

	http.SetCookie(res, &http.Cookie{
		Name:  "specific",
		Value: "some other values about specific things",
	})
	fmt.Fprintln(res, "COOKIES WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(res, "in chrome go to: dev tools/ applications / cookies")
}

//set multiple cookies
