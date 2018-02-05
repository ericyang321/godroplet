package main

import (
    "fmt"
    "github.com/ericyang321/godrople/util"
    "log"
    "net/http"
)

func infoHandler(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(res, "%s %s %s \n", req.Method, req.URL, req.Proto)

    for k, v := range req.Header {
        fmt.Fprintf(res, "Header[%q] = %q\n", k, v)
        fmt.Fprintf(res, "Host = %q\n", req.Host)
        fmt.Fprintf(res, "RemoteAddr = %q\n", req.RemoteAddr)
    }

    if err := req.ParseForm(); err != nil {
        log.Print(err)
    }

    for k, v := range req.Form {
        fmt.Fprintf(res, "Form[%q] = %q\n", k, v)
    }
}

func tempHandler(res http.ResponseWriter, req *http.Request) {

}

func main() {
    fmt.Println("Up and Running at :3000")
    http.HandleFunc("/", infoHandler)
    http.HandleFunc("/temperature", tempHandler)
    log.Fatal(http.ListenAndServe(":3000", nil))
}
