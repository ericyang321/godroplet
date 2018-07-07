package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	hn "github.com/ericyang321/godroplet/src/hn/client"
)

type templateData struct {
	Articles []hn.Article
}

func createHNHandler(num int, tpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var c hn.Client
		articles, err := c.GuaranteedTopArticles(num)
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
