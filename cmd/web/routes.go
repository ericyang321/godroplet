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
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileServer)).Methods("GET")
	router.HandleFunc("/", HandleHome).Methods("GET")
	router.Handle("/hn", CreateTimedHNHandler()).Methods("GET")

	return router
}
