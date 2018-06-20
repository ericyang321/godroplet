package main

import (
	"strings"

	"github.com/kr/pretty"
	"golang.org/x/net/html"
)

func main() {
	str := `<div className="woo"></div><div className="wah"></div>`
	n, _ := html.ParseFragment(strings.NewReader(str), &html.Node{
		Type: html.ElementNode,
	})

	pretty.Println(n)
	// html.Render(os.Stdout, n[0])
}
