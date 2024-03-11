package main

import (
	"Chi_Router_11/model"
	"database/sql"
	"log"
	"net/http"
	"text/template"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
)

var DBConn *sql.DB

func init() {
	var err error
	db, err := sql.Open("mysql", "root:Aniket@123@tcp(localhost:3306)/chi_htmx_demo")
	if err != nil {
		log.Println("Error in DB Connection", err)
	}
	DBConn = db
	log.Println("Database Connection Successfull")
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", homeHandler)
	r.Get("/user-info", userInfoHandler)
	r.Get("/posts", postHandler)

	http.ListenAndServe(":8080", r)
}

func homeHandler(res http.ResponseWriter, req *http.Request) {
	ctx := make(map[string]string)
	ctx["name"] = "Aniket"
	//Just for example that we can pass variables to template

	tpl, _ := template.ParseFiles("templates/index.html")
	err := tpl.Execute(res, ctx)

	if err != nil {
		log.Println("Error in template Execution", err.Error())
	}
}

func userInfoHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("User info from Api Server"))
}

func postHandler(res http.ResponseWriter, req *http.Request) {
	var posts []model.Post
	sql := "select * from posts"
	rows, err := DBConn.Query(sql)
	defer DBConn.Close()

	if err != nil {
		log.Println("error in DB execution", err)
	}

	for rows.Next() {
		data := model.Post{}

		err := rows.Scan(&data.Id, &data.Title)
		if err != nil {
			log.Println(err)
		}

		posts = append(posts, data)
	}

	log.Println(posts)
}
