package main

import (
	"log"
	"net/http"
	"runtime/debug"
)

// ServerError delivers error message, and leave stack trace in logs
func ServerError(w http.ResponseWriter, err error) {
	log.Printf("%s\n%s", err.Error(), debug.Stack())
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}
