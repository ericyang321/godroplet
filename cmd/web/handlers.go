package main

import (
	"encoding/json"
	"errors"
	"html/template"
	"net/http"
	"path/filepath"
	"time"

	"github.com/ericyang321/godroplet/pkg/hn"
	"github.com/ericyang321/godroplet/pkg/linkparser"
)

// returnErrJSON creates an error JSON and sends back bad request status code
func returnErrJSON(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	errInstance := linkparser.Error{Message: err.Error()}
	errJSON, _ := json.Marshal(errInstance)
	w.Write(errJSON)
}

// HandleLinkParser is a handler end point for user html submission
func HandleLinkParser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		returnErrJSON(w, errors.New("Needs to be a post request"))
		return
	}
	parseErr := r.ParseForm()
	if parseErr != nil {
		returnErrJSON(w, parseErr)
		return
	}
	links, extractErr := linkparser.Extract(r.Body)
	if extractErr != nil {
		returnErrJSON(w, extractErr)
		return
	}
	linkparser.LinksJSON(w, links)
}

// HandleHome is handler to serve personal site HTML
func HandleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	file := filepath.Join("./ui/html", "index.html")
	http.ServeFile(w, r, file)
}

// CreateTimedHNHandler initializes a single worker, which
// updates and fetches top 40 articles hacker news articles every 15 minutes.
func CreateTimedHNHandler() http.Handler {
	hnTemplate := template.Must(template.ParseFiles("./ui/html/hn.html"))
	return hn.CreateHNHandler(40, 15*time.Minute, hnTemplate)
}
