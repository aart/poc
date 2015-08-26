package main

import (
	"log"
	"net"
	"net/rpc"
)

import _ "github.com/go-sql-driver/mysql"

type Listener int




func (l *Listener) QueryData(in pack.Empty, out *[]pack.Person) error {
	personList := []pack.Person{ pack.Person{Name:"Aart",Firstname:"Verbeke"}, pack.Person{Name:"eeezzz",Firstname:"hhhzzz"}}
	log.Println("Query Data stub: ",personList)
	*out = personList
	return nil
}

func (l *Listener) SaveData(in pack.Person, ack *bool) error {
	log.Println("SaveData stub: ", in.Name, " ", in.Firstname)
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
