package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func redirectTLS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+r.Host+r.RequestURI, http.StatusMovedPermanently)
}

func main() {
	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", hello)
	log.Printf("Listening on %s ...\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
