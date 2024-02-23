package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8080", nil)
}

func myDog(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("COntent-Type", "text/html; charset=utf-8")
	io.WriteString(res, `<img src="photo4.jpg">`)
}
