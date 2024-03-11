package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	// Create a new Chi router
	r := chi.NewRouter()

	validCredentials := map[string]string{
		"user": "password",
	}

	// Middleware stack
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Define routes
	r.Get("/", homeHandler)
	r.Route("/users", func(r chi.Router) {
		r.Use(middleware.BasicAuth("user", validCredentials))
		r.Get("/", getUsers)
		r.Post("/", createUser)
		r.Get("/{userID}", getUser)
		r.Put("/{userID}", updateUser)
		r.Delete("/{userID}", deleteUser)
	})

	// Start the HTTP server
	http.ListenAndServe(":8080", r)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the home page!"))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List of users"))
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User created"))
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User details"))
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User updated"))
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User deleted"))
}
