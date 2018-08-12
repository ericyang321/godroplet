package hn

import (
	"html/template"
	"log"
	"net/http"
	"sync"
	"time"
)

type templateData struct {
	Articles []Article
}

// TODO: Separate worker's dependency on package fetch, and instead have Worker be pased in
// a fetcher interface

// Worker is a configurable periodic cache updater for scrapped hacker news articles.
// Work works in conjunction with package
type Worker struct {
	cache         *[]Article
	numOfArticles int
	tickDuration  time.Duration
	mutex         sync.Mutex
}

// InitializeTimer creates a timed instance that hot swaps cache of hacker news articles after certain intervals.
func (w *Worker) InitializeTimer() {
	var f Client
	go func() {
		tick := time.NewTicker(w.tickDuration)
		for {
			articles, err := f.GuaranteedTopArticles(w.numOfArticles)
			if err != nil {
				log.Fatal(err)
			}

			w.mutex.Lock()
			*(w.cache) = articles
			w.mutex.Unlock()

			<-tick.C
		}
	}()
}

// CreateHNHandler generates a convenient HTTP handler that
func CreateHNHandler(num int, duration time.Duration, tpl *template.Template) http.HandlerFunc {
	var cache []Article
	worker := Worker{
		numOfArticles: num,
		tickDuration:  duration,
		cache:         &cache,
	}
	worker.InitializeTimer()
	return func(w http.ResponseWriter, r *http.Request) {
		worker.mutex.Lock()
		tmpldata := templateData{Articles: cache}
		worker.mutex.Unlock()

		err := tpl.Execute(w, tmpldata)
		if err != nil {
			log.Fatal(err)
		}
	}
}
