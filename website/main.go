package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("statics/templates/index.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, "")
}

func setStaticPath() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("statics"))))
}

func initializeTemplate(template *template.Template) *template.Template {
	contents, err := ioutil.ReadFile(string(template.Name() + ".html.tmpl"))
	if err != nil {
		log.Panic(err)
	}
	template.Parse(string(contents))
	return template
}

func main() {
	setStaticPath()
	// layout := initializeTemplate(template.New("layout"))
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
