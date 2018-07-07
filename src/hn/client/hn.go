package hn

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// API_URL is hacker new's official firebase url for querying information
const API_URL = "https://hacker-news.firebaseio.com/v0"

// YCOMB_ITEM_URL is hacker new's official website url for accessing specific items (comment, article, etc)
const YCOMB_ITEM_URL = "https://news.ycombinator.com/item?id="

// Client is http wrapper that fetches hacker news data
type Client struct {
	apiURL string
}

// Item represents a single item returned by the HN API
// for this hacker HN website, we only care about stories and jobs
type Item struct {
	By    string `json:"by"`
	ID    int    `json:"id"`
	Score int    `json:"score"`
	Time  int    `json:"time"`
	Title string `json:"title"`
	Type  string `json:"type"`
	Text  string `json:"text"`
	URL   string `json:"url"`
	// Unused keys
	// Kids        []int  `json:"kids"`
	// Descendants int    `json:"descendants"`
}

// Article is decoded hacker news JSON structure (icon) along with parsed article host
// and hacker news post url
type Article struct {
	Item
	Host     string
	YcombURL string
}

// pylon is used to relay successfully decoded Item or error through go channel
type pylon struct {
	item Item
	err  error
}

func (c *Client) init() {
	if c.apiURL == "" {
		c.apiURL = API_URL
	}
}

// GetTopIds fetches a list of top 100 hacker news ID articles, and save them on passed list
func (c *Client) GetTopIds() ([]int, error) {
	c.init()

	var topIDs []int
	res, getErr := http.Get(c.apiURL + "/topstories.json")
	if getErr != nil {
		return topIDs, getErr
	}
	defer res.Body.Close()
	decodeErr := json.NewDecoder(res.Body).Decode(&topIDs)
	if decodeErr != nil {
		return topIDs, decodeErr
	}
	return topIDs, nil
}

// GetItem http fetch a single hacker news item by its id
func (c *Client) GetItem(id int) (Item, error) {
	c.init()

	var item Item
	url := fmt.Sprintf("%s/item/%d.json", c.apiURL, id)
	res, getErr := http.Get(url)
	if getErr != nil {
		return item, getErr
	}
	defer res.Body.Close()
	decodeErr := json.NewDecoder(res.Body).Decode(&item)
	if decodeErr != nil {
		return item, decodeErr
	}

	return item, nil
}

// GetArticles fetches top num of articles
func (c *Client) GetArticles(num int) ([]Article, error) {
	c.init()
	var articles []Article
	topIDs, err := c.GetTopIds()
	if err != nil {
		return nil, err
	}
	for _, id := range topIDs {
		receiver := make(chan pylon)
		go c.asyncFetch(id, receiver)

		py := <-receiver

		if py.err != nil {
			continue
		} else {
			article := createArticle(py.item)
			articles = append(articles, article)
		}
		if len(articles) >= num {
			break
		}
	}
	return articles, nil
}

func (c *Client) asyncFetch(id int, receiver chan pylon) {
	c.init()
	item, err := c.GetItem(id)
	if err != nil {
		receiver <- pylon{err: err}
		return
	}
	receiver <- pylon{item: item}
}

func isStory(article Article) bool {
	return article.Type == "story" && article.URL != ""
}

func createArticle(item Item) Article {
	article := Article{Item: item}
	articleURL, err := url.Parse(article.URL)
	if err == nil {
		article.Host = strings.TrimPrefix(articleURL.Hostname(), "www.")
		article.YcombURL = fmt.Sprintf("%s%d", YCOMB_ITEM_URL, article.ID)
	}
	return article
}
