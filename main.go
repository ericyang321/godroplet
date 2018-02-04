package main

import (
    "fmt"
    "net/http"
)

func handler(res http.ResponseWriter, req *http.Request) {
    res.WriteHeader(200)
    res.Write([]byte("OK"))
}

func main() {
    fmt.Println("Up and Running at :3000")
    http.HandleFunc("/", handler)
    http.ListenAndServe(":3000", nil)
}
