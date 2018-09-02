package main

import (
	"bytes"
	"html/template"
	"net/http"
	"path/filepath"
	"time"

	"github.com/ericyang321/godroplet/pkg/hn"
)

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
	return createHNHandler(40, 15*time.Minute, hnTemplate)
}

// createHNHandler generates a convenient HTTP handler that
func createHNHandler(num int, duration time.Duration, tpl *template.Template) http.HandlerFunc {
	var cache []hn.Article
	worker := hn.Worker{
		NumOfArticles: num,
		TickDuration:  duration,
		Cache:         &cache,
	}
	worker.InitializeTimer()
	return func(w http.ResponseWriter, r *http.Request) {
		worker.Mutex.Lock()
		tmpldata := hn.TemplateData{Articles: cache}
		worker.Mutex.Unlock()

		buf := new(bytes.Buffer)
		err := tpl.Execute(buf, tmpldata)
		if err != nil {
			ServerError(w, err)
			return
		}
		buf.WriteTo(w)
	}
}
