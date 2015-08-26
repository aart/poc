package main

import _ "github.com/go-sql-driver/mysql"
import (
	"database/sql"
	"fmt"
	"log"
)

func testDB() {
	db, err := sql.Open("mysql", "root:@tcp(10.240.61.254:3306)/test")
	if err != nil {
		log.Fatal(err) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM person")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for rows.Next() {

		var name string
		var firstname string
		// get RawBytes from data
		err = rows.Scan(&name, &firstname)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		fmt.Println(name, " ", firstname)

	}
}

func main() {

	fmt.Println("querying the DB")
	testDB()

}
