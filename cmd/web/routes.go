package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// CreateMux mounts routes and sets up static file urls
func CreateMux() *mux.Router {
	router := mux.NewRouter()
	fileServer := http.FileServer(http.Dir("./ui/static"))

	// Routes
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileServer))
	router.HandleFunc("/", HandleHome)

	router.Handle("/hn", CreateTimedHNHandler())

	return router
}
