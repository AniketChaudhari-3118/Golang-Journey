package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	ID          int
	Name        string
	Price       float32
	Description string
}

var tpl *template.Template
var db *sql.DB

func main() {
	tpl, _ = template.ParseGlob("template/*")
	var err error
	db, err = sql.Open("mysql", "root:Aniket@123@tcp(localhost:3306)/testdb")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	http.HandleFunc("/insert", insertHandler)
	http.HandleFunc("/browse", browseHandler)
	http.HandleFunc("/update/", updateHandler)
	http.HandleFunc("/updaterResult/", updaterResultHandler)
	http.HandleFunc("/delete/", deleteHandler)
	http.HandleFunc("/", homePageHandler)
	http.ListenAndServe(":8080", nil)
}

func browseHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("*****browseHandler Running*****")
	stmt := "select * from products"
	rows, err := db.Query(stmt)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price, &p.Description)
		if err != nil {
			panic(err)
		}
		products = append(products, p)
	}

	tpl.ExecuteTemplate(res, "select.html", products)
}

func insertHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("*****insertHandler Running*****")
	if req.Method == "GET" {
		tpl.ExecuteTemplate(res, "insert.html", nil)
		return
	}
	req.ParseForm()

	id := req.FormValue("id")
	name := req.FormValue("nameName")
	price := req.FormValue("priceName")
	descr := req.FormValue("descrName")
	var err error
	if name == "" || price == "" || descr == "" || id == "" {
		fmt.Println("Error inserting row: ", err)
		tpl.ExecuteTemplate(res, "insert.html", "error inserting data, please check all feilds")
	}

	var ins *sql.Stmt
	ins, err = db.Prepare("insert into `testdb`.`products` (`id`, `name`, `price`, `description`) values (?, ?, ?, ?);")
	if err != nil {
		panic(err)
	}
	defer ins.Close()

	resu, err := ins.Exec(id, name, price, descr)
	rowsAffec, _ := resu.RowsAffected()
	if err != nil || rowsAffec != 1 {
		fmt.Println("error inserting row:", err)
		tpl.ExecuteTemplate(res, "insert.html", "Error inserting data, please check all fields ")
		return
	}

	lastInserted, _ := resu.LastInsertId()
	rowsAffected, _ := resu.RowsAffected()
	fmt.Println("ID of last row inserted:", lastInserted)
	fmt.Println("number of rows affected:", rowsAffected)
	tpl.ExecuteTemplate(res, "insert.html", "Product inserted Successfully")
}

func updateHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("*****updateHandler Running*****")
	req.ParseForm()
	id := req.FormValue("idproducts")
	row := db.QueryRow("select * from testdb.products where idproducts = ?;", id)
	var p Product
	err := row.Scan(&p.ID, &p.Name, &p.Price, &p.Description)
	if err != nil {
		fmt.Println(err)
		http.Redirect(res, req, "/browse", http.StatusTemporaryRedirect)
		return
	}
	tpl.ExecuteTemplate(res, "update.html", p)
}

func updaterResultHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("*****updateResultHandler Running*****")
	req.ParseForm()
	id := req.FormValue("idProducts")
	name := req.FormValue("nameName")
	price := req.FormValue("priceName")
	description := req.FormValue("descrName")
	upStmt := "update `testdb`.`products` set `name` = ?, `price` = ?, `description` = ? where (`id` = ?);"
	stmt, err := db.Prepare(upStmt)
	if err != nil {
		fmt.Println("err preparing statement")
		panic(err)
	}
	fmt.Println("db.Prepare err: ", err)
	fmt.Println("db.Prepare stmt: ", stmt)
	defer stmt.Close()
	var resu sql.Result
	resu, err = stmt.Exec(name, price, description, id)
	rowsAff, _ := resu.RowsAffected()
	if err != nil || rowsAff != 1 {
		fmt.Println(err)
		tpl.ExecuteTemplate(res, "result.html", "their was a problem updating the product")
		return
	}
	tpl.ExecuteTemplate(res, "result.html", "Product was Successfully Updated")

}

func deleteHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("*****deleteHandler Running*****")
	req.ParseForm()
	id := req.FormValue("idproducts")
	del, err := db.Prepare("DELETE FROM `testdb`.`products` WHERE (`id` = ?);")
	if err != nil {
		panic(err)
	}
	defer del.Close()
	var resu sql.Result
	resu, err = del.Exec(id)
	rowsAff, _ := resu.RowsAffected()
	fmt.Println("rowsAff: ", rowsAff)

	if err != nil || rowsAff != 1 {
		fmt.Fprint(res, "error deleting product")
	}
	fmt.Println("err:", err)
	tpl.ExecuteTemplate(res, "result.html", "Product was Successfully Deleted")
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/browse", http.StatusTemporaryRedirect)
}
