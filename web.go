package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/rpc"
	"poc/pack"
)

func heartbeat(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "heartbeat")
	log.Println("web layer - heartbeat")

}

func queryData(w http.ResponseWriter, r *http.Request) {

	client, err := rpc.Dial("tcp", "localhost:42586")
	if err != nil {
		log.Fatal(err)
	}

	var out []pack.Person
	in := pack.Empty{}

	err = client.Call("Listener.QueryData", in, &out)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("web layer - queryData: ", out)

	b, err := json.Marshal(out)
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(w, string(b))

}

func saveData(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var in pack.Person
	err := decoder.Decode(&in)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("web layer - saveData: ", in.Name, " ", in.Firstname)

	client, err := rpc.Dial("tcp", "localhost:42586")
	if err != nil {
		log.Fatal(err)
	}

	var out bool

	err = client.Call("Listener.SaveData", in, &out)
	if err != nil {
		log.Fatal(err)
	}

}

func getTransportOrder(w http.ResponseWriter, r *http.Request) {

	u := r.URL
	q := u.Query()
	in := q.Get("id")

	client, err := rpc.Dial("tcp", "localhost:42586")
	if err != nil {
		log.Fatal(err)
	}

	var out pack.TransportOrder

	err = client.Call("Listener.GetTransportOrderById", in, &out)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("web layer - getTransportOrder: ", out)

	b, err := json.MarshalIndent(out, "", "   ")
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(w, string(b))

}

func main() {

	http.HandleFunc("/", heartbeat)
	http.HandleFunc("/queryData", queryData)
	http.HandleFunc("/getTransportOrder", getTransportOrder)
	http.HandleFunc("/saveData", saveData)
	http.ListenAndServe(":8080", nil)
}
