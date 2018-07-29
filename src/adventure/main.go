package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type arc struct {
	Title   string
	Story   []string
	Options []option
}

type option struct {
	Text string
	Arc  string
}

func createHandler(key string, blob *map[string]arc) http.HandlerFunc {
	template := template.Must(template.ParseFiles("./template.html"))
	chosenArc := (*blob)[key]
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		if err := template.Execute(w, chosenArc); err != nil {
			log.Fatal(err)
		}
	}
}

func slurpJSON() map[string]arc {
	raw, readErr := ioutil.ReadFile("./story.json")
	blob := make(map[string]arc)
	if readErr != nil {
		fmt.Println(readErr.Error())
		os.Exit(1)
	}
	jsonErr := json.Unmarshal(raw, &blob)
	if jsonErr != nil {
		fmt.Println(jsonErr.Error())
		os.Exit(1)
	}
	return blob
}

func main() {
	blob := slurpJSON()
	mux := http.NewServeMux()
	for key := range blob {
		mux.HandleFunc("/"+key, createHandler(key, &blob))
	}
	http.ListenAndServe("localhost:8080", mux)
}
