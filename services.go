package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"database/sql"
)

import _ "github.com/go-sql-driver/mysql"

type Listener int

type Empty struct {
}

type Person struct {
	Name      string
	Firstname string
}

func (l *Listener) QueryData(in Empty, out *[]Person) error {
	*out = queryAllPersonsDB()
	return nil
}

func (l *Listener) SaveData(in Person, ack *bool) error {
	fmt.Println(in.Name, " ", in.Firstname)
	return nil
}

func queryAllPersonsDB() []Person {
	db, err := sql.Open("mysql", "root:@tcp(10.240.61.254:3306)/test")
	if err != nil {
		log.Fatal(err) 
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM person")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	personList := make([]Person, 10)

	for rows.Next() {

		var name string
		var firstname string
		// get RawBytes from data
		err = rows.Scan(&name, &firstname)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		p := Person{Name: name, Firstname: firstname}
		personList = append(personList, p)

	}

	return personList
}

func main() {

	addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0:42586")
	if err != nil {
		log.Fatal(err)
	}

	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		log.Fatal(err)
	}

	listener := new(Listener)
	rpc.Register(listener)
	rpc.Accept(inbound)
}
