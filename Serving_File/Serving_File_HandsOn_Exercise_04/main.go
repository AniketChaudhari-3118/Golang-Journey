package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/resource", http.StripPrefix("resources", http.FileServer(http.Dir("./assets"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}


// Serve the files in the "starting-files" folder
// To get your images to serve, use:
// 	func StripPrefix(prefix string, h Handler) Handler
// 	func FileServer(root FileSystem) Handler
// Constraint: you are not allowed to change the route being used for images in the template file