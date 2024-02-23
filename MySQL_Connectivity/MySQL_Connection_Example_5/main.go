package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var tpl *template.Template
var db *sql.DB

func main() {
	tpl, _ = template.ParseGlob("template/*")
	var err error
	db, err = sql.Open("mysql", "root:Aniket@123@tcp(localhost:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/loginauth", loginAuthHandler)
	http.ListenAndServe(":8080", nil)
}

func loginHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("*****loginHandler Running*****")
	tpl.ExecuteTemplate(res, "login.html", nil)
}

func loginAuthHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("*****loginAuthHandler Running*****")
	req.ParseForm()
	user_name := req.FormValue("username")
	password := req.FormValue("password")

	// aniket_password, err := hasPassword("Aniket@123")
	password_new := "Aniket_chaudhari"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password_new), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	// upStmt := "update `testdb`.`users` set `password` = hashedPassword,  where (`username` = Aniket);"
	stmt, err := db.Prepare("INSERT INTO users (username, hashed_password) VALUES (?, ?)")
	if err != nil {
		fmt.Println(err.Error())
	}
	var resu sql.Result
	resu, err = stmt.Exec(user_name, hashedPassword)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(resu)

	var value string
	row := db.QueryRow("SELECT hashed_password FROM users WHERE username = ?", user_name)
	err = row.Scan(&value)
	if err != nil {
		panic(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(value), []byte(password))
	if err == nil {
		fmt.Fprintf(res, "You have Successfully Logged in")
		cookie := &http.Cookie{
			Name:     user_name,
			Value:    password,
			Path:     "/",
			HttpOnly: true,
			MaxAge:   -1,
		}
		http.SetCookie(res, cookie)
		return
	} else {
		fmt.Println("Incorrect Password")
		tpl.ExecuteTemplate(res, "login.html", "check username and password")
	}
}
