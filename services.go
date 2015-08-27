package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"poc/pack"
)

import _ "github.com/go-sql-driver/mysql"

type Listener int

func (l *Listener) QueryData(in pack.Empty, out *[]pack.Person) error {
	*out = queryAllPersonsDB()
	return nil
}

func (l *Listener) SaveData(in pack.Person, ack *bool) error {
	fmt.Println(in.Name, " ", in.Firstname)

	// extract as function
	db, err := sql.Open("mysql", "root:@tcp(10.240.61.254:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
    
    stmtIns, err := db.Prepare("INSERT INTO persons VALUES( ?, ? )") // ? = placeholder
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    defer stmtIns.Close()

	_, err = stmtIns.Exec(in.Name, in.Firstname)


	return nil
}




func queryAllPersonsDB() []pack.Person {
	db, err := sql.Open("mysql", "root:@tcp(10.240.61.254:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM person")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	personList := make([]pack.Person,0)

	for rows.Next() {

		var name string
		var firstname string
		// get RawBytes from data
		err = rows.Scan(&name, &firstname)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		p := pack.Person{Name: name, Firstname: firstname}
		personList = append(personList, p)

	}

	return personList
}


func (l *Listener) GetTransportOrderById(in string, out *pack.TransportOrder) error {

	db, err := sql.Open("mysql", "root:@tcp(10.240.61.254:3306)/test")
	
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rowA := db.QueryRow("SELECT * FROM transportorder WHERE businessid = ?", in)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var BusinessId  string
	var Carrier     string
	var Express     bool
	var ContractRef string
	var Goods       string
	var Origin      string
	var Destination string

	err = rowA.Scan(&BusinessId,&Carrier,&Express,&ContractRef,&Goods,&Origin,&Destination)
	if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
	}

	rowB := db.QueryRow("SELECT * FROM goods WHERE id = ?", Goods)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var g pack.Goods

	err = rowB.Scan(&g.Id,&g.Description,&g.Bulk,&g.TotalLoading,&g.TotalNetWeight,&g.TotalVolume,&g.TotalPackage,&g.TotalPallets)

	rowC := db.QueryRow("SELECT * FROM endpoint WHERE id = ?", Origin)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var o pack.Endpoint

	err = rowC.Scan(&o.Id,&o.Detail)

	rowD := db.QueryRow("SELECT * FROM endpoint WHERE businessdd = ?", Destination)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var d pack.Endpoint

	err = rowD.Scan(&d.Id,&d.Detail)

to := pack.TransportOrder{
		BusinessId:  BusinessId,
		Carrier:     Carrier,
		Express:    Express,
		ContractRef: ContractRef,
		Goods: g,
		Origin: o,
		Destination: d,
	}


	*out = to

	return nil
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
