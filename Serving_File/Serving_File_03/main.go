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
}

func myPic(res http.ResponseWriter, req *http.Request) {
	f, err := os.Open("photo4.jpg")
	if err != nil {
		http.Error(res, "File not found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(res, "file not found", 404)
		return
	}

	http.ServeContent(res, req, f.Name(), fi.ModTime(), f)
}
