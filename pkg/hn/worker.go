package hn

import (
	"log"
	"sync"
	"time"
)

// TemplateData is struct specifically for holding template parsing data
type TemplateData struct {
	Articles []Article
}

// Worker is a configurable periodic cache updater for scrapped hacker news articles.
// works in conjunction with fetcher
type Worker struct {
	Cache         *[]Article
	NumOfArticles int
	TickDuration  time.Duration
	Mutex         sync.Mutex
}

// InitializeTimer creates a timed instance that hot swaps cache of hacker news articles after certain intervals.
func (w *Worker) InitializeTimer() {
	var c Client
	go func() {
		tick := time.NewTicker(w.TickDuration)
		for {
			articles, err := c.GuaranteedTopArticles(w.NumOfArticles)
			if err != nil {
				log.Fatal(err)
			}

			w.Mutex.Lock()
			*(w.Cache) = articles
			w.Mutex.Unlock()

			<-tick.C
		}
	}()
}
