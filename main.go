package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(res http.ResponseWriter, req *http.Request) {
    // Echo HTTPS Request
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

func main() {
    fmt.Println("Up and Running at :3000")
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":3000", nil))
}
