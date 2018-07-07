package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	hn "github.com/ericyang321/godroplet/src/hn/client"
)

var cache []hn.Article
var cacheExpiration time.Time

type templateData struct {
	Articles []hn.Article
}

func getArticles(num int) ([]hn.Article, error) {
	if time.Now().Sub(cacheExpiration) < 0 {
		return cache, nil
	}
	var c hn.Client
	articles, err := c.GuaranteedTopArticles(num)
	if err != nil {
		return nil, err
	}
	cache = articles
	cacheExpiration = time.Now().Add(1 * time.Second)
	return cache, nil
}

func createHNHandler(num int, tpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		articles, err := getArticles(num)
		if err != nil {
			handleFailure(w, err)
			return
		}
		tmpldata := templateData{Articles: articles}
		err = tpl.Execute(w, tmpldata)
		if err != nil {
			handleFailure(w, err)
			return
		}
	}
}

func handleFailure(w http.ResponseWriter, err error) {
	fmt.Fprintf(w, "You fucked up real bad: \n"+err.Error())
}

func main() {
	tpl := template.Must(template.ParseFiles("./index.html"))
	http.HandleFunc("/", createHNHandler(30, tpl))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 3000), nil))
}
