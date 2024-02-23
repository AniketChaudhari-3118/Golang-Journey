package main

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		id := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			//Secure: true,
			HttpOnly: true, // it cannot be accessed by javascript
		}
		http.SetCookie(res, cookie)
	}
	fmt.Println(cookie)
}
