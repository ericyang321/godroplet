package main

import (
	"strings"

	"github.com/ericyang321/godroplet/src/trim/helper"
	"github.com/kr/pretty"
	"golang.org/x/net/html"
)

func main() {
	text := `
	<modal class="wah">
		Hello its me a modal
	</modal>
	`
	n, _ := html.Parse(strings.NewReader(text))
	list := helper.DeepGrep(n, func(n *html.Node) bool {
		return n.Data == "modal"
	})

	for _, np := range list {
		n := *np
		pretty.Println("Name: ", n.Data)
		pretty.Println("Attributes: ", n.Attr)
	}
}
