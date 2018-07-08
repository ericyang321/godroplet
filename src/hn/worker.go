package worker

import (
	"html/template"
	"log"
	"net/http"
	"sync"
	"time"

	fetch "github.com/ericyang321/godroplet/src/hn/client"
)

type templateData struct {
	Articles []fetch.Article
}

// TODO: Separate worker's dependency on package fetch, and instead have Worker be pased in
// a fetcher interface

// Worker is a configurable periodic cache updater for scrapped hacker news articles.
// Work works in conjunction with package fetch.
type Worker struct {
	cache         *[]fetch.Article
	numOfArticles int
	tickDuration  time.Duration
	mutex         sync.Mutex
}

// InitializeTimer creates a timed instance that hot swaps cache of hacker news articles after certain intervals.
func (w *Worker) InitializeTimer() {
	var f fetch.Client
	go func() {
		tick := time.NewTicker(w.tickDuration)
		for {
			articles, err := f.GuaranteedTopArticles(w.numOfArticles)
			if err != nil {
				printErr(err.Error())
			}

			w.mutex.Lock()
			*(w.cache) = articles
			w.mutex.Unlock()

			<-tick.C
		}
	}()
}

// CreateHNHandler generates a convenient HTTP handler that initializes a single worker, which
// updates and fetches top 30 articles hacker news articles every 15 minutes.
func CreateHNHandler(num int, duration time.Duration, tpl *template.Template) http.HandlerFunc {
	var cache []fetch.Article
	worker := Worker{
		numOfArticles: 40,
		tickDuration:  15 * time.Minute,
		cache:         &cache,
	}
	worker.InitializeTimer()
	return func(w http.ResponseWriter, r *http.Request) {
		worker.mutex.Lock()
		tmpldata := templateData{Articles: cache}
		worker.mutex.Unlock()

		err := tpl.Execute(w, tmpldata)
		if err != nil {
			printErr(err.Error())
		}
	}
}

func printErr(err string) {
	log.Printf("===== ERROR =====")
	log.Printf(err)
}
