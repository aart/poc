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


func (l *Listener) GetTransportOrder(in pack.Empty, out *pack.TransportOrder) error {
	
to := pack.TransportOrder{
		BusinessId:  "8678900",
		Carrier:     "ABCLogistics",
		Express:     false,
		ContractRef: "5678890DDC",
		Goods: pack.Goods{
			Id:             "6543457898",
			Description:    "fine goods",
			Bulk:           false,
			TotalLoading:   122,
			TotalNetWeight: 6788,
			TotalVolume:    5678,
			TotalPackage:   89900,
			TotalPallets:   778889,
		},
		Origin: pack.Endpoint{
            Detail: "Endpoint A",
        },
		Destination: pack.Endpoint{
			Detail: "Endpoint B",
		},
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
