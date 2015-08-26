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

	log.Println("queryData: ",out)

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
	log.Println("saveData: ",in.Name, " ", in.Firstname)

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

func main() {

	http.HandleFunc("/", heartbeat)
	http.HandleFunc("/queryData", queryData)
	http.HandleFunc("/saveData", saveData)
	http.ListenAndServe(":8080", nil)
}
