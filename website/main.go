package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func getTemplatePath(name string) string {
	return filepath.Join("statics", "templates", name)
}

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

func main() {
	setStaticPath()
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
