package main

import (
    "net/http"
)

func handler(res http.ResponseWriter, req *http.Request) {
    res.WriteHeader(200)
    res.Write([]byte("OK"))
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":3000", nil)
}
