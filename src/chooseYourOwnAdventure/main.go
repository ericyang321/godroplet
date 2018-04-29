package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func buildTemplate(key string, blob map[string]interface{}) {

}

func createHandler(key string, blob map[string]interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func slurpJSON() map[string]interface{} {
	raw, readErr := ioutil.ReadFile("./gopher.json")
	blob := make(map[string]interface{})
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
	for k := range blob {
		mux.HandleFunc("/"+k, createHandler(k, blob))
	}
	http.ListenAndServe("localhost:8080", mux)
}
