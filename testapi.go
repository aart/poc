package main

import (
	"bytes"
	//"encoding/json"
	"fmt"
	"io/ioutil"
	//"log"
	"net/http"
)

func testHeartbeat() {
	url := "http://localhost:8080"
	fmt.Println("URL:>", url)

	req, err := http.NewRequest("POST", url, nil)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

//------------------------------------------------------
func testSaveData() {
	url := "http://localhost:8080/saveData"
	fmt.Println("URL:>", url)

	var jsonStr = []byte(`{"Name":"Verbeeke", "Firstname":"Bart"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

//------------------------------------------------------
func testQueryData() {
	url := "http://localhost:8080/queryData"
	fmt.Println("URL:>", url)

	req, err := http.NewRequest("POST", url, nil)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

//------------------------------------------------------
// test nested structs
func testGetTransportOrder() {
    url := "http://localhost:8080/getTransportOrder?id=1"
    fmt.Println("URL:>", url)

    req, err := http.NewRequest("POST", url, nil)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}


func main() {
	testGetTransportOrder()
	testHeartbeat()
	testSaveData()
	testQueryData()
}
