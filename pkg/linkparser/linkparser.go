package linkparser

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// Link represents an <a> link in an HTML document
type Link struct {
	Href string
	Text string
}

// Error represents any return messages of faulty HTML traversing
type Error struct {
	Message string
}

// Links represent end return json for user after links are parsed
type Links struct {
	Result []Link
}

// LinksJSON
func LinksJSON(w http.ResponseWriter, linksList []Link) {
	w.WriteHeader(http.StatusOK)
	linksInstance := Links{Result: linksList}
	linksJSON, _ := json.Marshal(linksInstance)
	w.Write(linksJSON)
}

// createNodesList depth first recursively traverses generated html tree
// and pushes <a> tags into slice.
func createNodesList(n *html.Node) []*html.Node {
	// base case
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var nodeList []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodeList = append(nodeList, createNodesList(c)...)
	}
	return nodeList
}

func instantiateLink(n *html.Node) Link {
	newLink := Link{}
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			newLink.Href = attr.Key
		}
		newLink.Text = getAllText(n)
	}
	return newLink
}

// <a> tags may have text inside surrounded by other tags.
// This function strips the tags and accumulate all text nodes
// inside a given <a> tag
func getAllText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	// ignore comment nodes:
	if n.Type != html.ElementNode {
		return ""
	}
	str := ""
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		str += getAllText(c) + " "
	}

	return clean(str)
}

func clean(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func createLinksList(nodeList []*html.Node) []Link {
	var linksList []Link
	for _, node := range nodeList {
		linksList = append(linksList, instantiateLink(node))
	}
	return linksList
}

// Extract converts HTML document from reader and return slice of links
func Extract(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodeList := createNodesList(doc)
	linkList := createLinksList(nodeList)

	return linkList, nil
}
