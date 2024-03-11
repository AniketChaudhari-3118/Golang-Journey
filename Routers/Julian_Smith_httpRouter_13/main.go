package main

import (
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/protected", Middleware(protectedHandler))

	router.GET("/", Logger(IndexHandler))
	router.GET("/hello/:name", Logger(HelloHandler))

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}

func Logger(h httprouter.Handle) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		//Log the request method and url path
		log.Printf("%s %s", req.Method, req.URL.Path)

		//call the next Handler
		h(res, req, ps)
	}
}

func Recovery(h http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		//Defer a function to recover from any panics
		defer func() {
			if err := recover(); err != nil {
				//Log the panic and respond with an internal server error
				log.Printf("Panic: %v", err)
				http.Error(res, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		h.ServeHTTP(res, req)
	})
}

func Middleware(next httprouter.Handle) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		// Middleware Logic goes here

		log.Println("Executing Middleware....")
		//Call the next Handler
		next(res, req, ps)
	}
}

func protectedHandler(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	res.Write([]byte("This is protected resource"))
}

func IndexHandler(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	res.Write([]byte("Welcome to the HomePage!"))
	log.Fatalln()
}

func HelloHandler(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	res.Write([]byte("Hello, " + name + "!"))
}
