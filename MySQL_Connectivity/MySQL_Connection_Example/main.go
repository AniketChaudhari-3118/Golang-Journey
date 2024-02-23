package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Open may just validate its arguments without creating a connection to the database.
	// To verify that the data source name is valid, call ping.
	//func open(driverName, dataSourceName string) (*DB, error)
	//dataSourceName: username:password@protocol(address)/dbname?param=value
	db, err := sql.Open("mysql", "root:Aniket@123@tcp(localhost:3306)/testdb")
	if err != nil {
		fmt.Println("error validating sql.Open arguments")
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("error verifying connection with db.Ping")
		panic(err.Error())
	}

	//func (db *DB) Query(query string, args ...interdace()) (*Rows, error)
	insert, err := db.Query("insert into `testdb`.`test`(`Id`, `Firstname`, `Lastname`) values ('6', 'Lionel', 'Messi');")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Successful Connection to Database! ")
}
