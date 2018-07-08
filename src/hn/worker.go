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

// Worker is
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
				log.Printf("============ ERROR: ===========")
				log.Printf(err.Error())
			}

			w.mutex.Lock()
			*(w.cache) = articles
			w.mutex.Unlock()

			<-tick.C
		}
	}()
}

// CreateHNHandler generates HTTP handler
func CreateHNHandler(num int, duration time.Duration, tpl *template.Template) http.HandlerFunc {
	var cache []fetch.Article
	worker := Worker{
		numOfArticles: 30,
		tickDuration:  1 * time.Second,
		cache:         &cache,
	}
	worker.InitializeTimer()
	return func(w http.ResponseWriter, r *http.Request) {
		worker.mutex.Lock()
		tmpldata := templateData{Articles: cache}
		worker.mutex.Unlock()

		err := tpl.Execute(w, tmpldata)
		if err != nil {
			log.Printf(" ========= Template error ======= \n" + err.Error())
			return
		}
	}
}
