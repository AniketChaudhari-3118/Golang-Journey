package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", me)
	http.HandleFunc("/photo4.jpg", myPic)
	http.ListenAndServe(":8082", nil)
}

func me(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `<img src="photo4.jpg">`)
	// io.WriteString(res, "hi its me")
}

func myPic(res http.ResponseWriter, h *http.Request) {
	f, err := os.Open("photo4.jpg")
	if err != nil {
		http.Error(res, "File not found", 404)
		return
	}
	defer f.Close()

	io.Copy(res, f)
}
