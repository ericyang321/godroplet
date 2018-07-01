package main

import (
	"encoding/json"
	"net/http"

	"github.com/kr/pretty"
)

const apiURL = "https://hacker-news.firebaseio.com/v0"

func getTopStories(topIDs *[]int) error {
	res, getErr := http.Get(apiURL + "/topstories.json")
	if getErr != nil {
		return getErr
	}
	defer res.Body.Close()
	decodeErr := json.NewDecoder(res.Body).Decode(topIDs)
	if decodeErr != nil {
		return decodeErr
	}
	return nil
}

func filterThirty(topIDs []int) []int {
	if len(topIDs) <= 30 {
		return topIDs
	}
	onlyThirtyIds := make([]int, 30)
	for i := 0; i < 30; i++ {
		onlyThirtyIds[i] = topIDs[i]
	}
	return onlyThirtyIds
}

func main() {
	topIDs := make([]int, 100)
	fetchErr := getTopStories(&topIDs)
	if fetchErr != nil {
		panic(fetchErr.Error())
	}
	topThirtyIDs := filterThirty(topIDs)
	pretty.Println(topThirtyIDs)
}
