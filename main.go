package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ericyang321/godroplet/src/linkparser"
)

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	env := os.Getenv("ENV")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	if env == "development" {
		return "localhost:" + port, nil
	}
	return ":" + port, nil
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
	// Routes
	mux.Handle("/", http.FileServer(http.Dir("./src/assets")))
	mux.HandleFunc("/parse-link-tags", linkparser.HandlerFunc)

	// Force HTTPS redirect
	secureMux := redirectTLS(mux)
	log.Printf("Listening on %s ...\n", addr)
	log.Fatal(http.ListenAndServe(addr, secureMux))
}
