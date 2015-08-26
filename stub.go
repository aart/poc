package main

import (
	"log"
	"net"
	"net/rpc"
	"poc/pack"
)

import _ "github.com/go-sql-driver/mysql"

type Listener int

func (l *Listener) QueryData(in pack.Empty, out *[]pack.Person) error {
	personList := []pack.Person{pack.Person{Name: "Aart", Firstname: "Verbeke"}, pack.Person{Name: "Katrien", Firstname: "De Muynck"}}
	log.Println("Query Data stub: ", personList)
	*out = personList
	return nil
}

func (l *Listener) SaveData(in pack.Person, ack *bool) error {
	log.Println("SaveData stub: ", in.Name, " ", in.Firstname)
	return nil
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
