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

func redirectTLS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		proto := r.Header.Get("x-forwarded-proto")
		if proto == "http" || proto == "HTTP" {
			http.Redirect(w, r, fmt.Sprintf("https://%s%s", r.Host, r.URL), http.StatusPermanentRedirect)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

	secureMux := redirectTLS(mux)
	log.Printf("Listening on %s ...\n", addr)
	log.Fatal(http.ListenAndServe(addr, secureMux))
}
