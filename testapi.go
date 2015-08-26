package main

import (
  "bytes"
  "fmt"
  "io/ioutil"
  "net/http"
  "encoding/json"
  "log"
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

type Goods struct {
    Name            string
    Description    string
}

type Det struct{
    Detail string
}

type TransportOrder struct {
    Id string
    Gd Goods
    Det
}


func getData() {
    
    to := TransportOrder{
            Id : "businessID",
            Gd : Goods{
                Name : "pallets",
                Description : "blabla",
            },
            Det : Det{Detail : "detail...."},
    }
    fmt.Println(to.Detail)

     fmt.Println(to)

    jsonStr, err := json.Marshal(to)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(jsonStr))
}


func main() {
    testHeartbeat()
    testSaveData()
    testQueryData()
    getData()
}
