package helper

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Matcher is user passable function to decide which nodes to fetch
type Matcher func(*html.Node) bool

// IsPhrasingContent Determine if a node qualifies as phrasing content.
// https://developer.mozilla.org/en-US/docs/Web/Guide/HTML/Content_categories#Phrasing_content
func IsPhrasingContent(n *html.Node) bool {
	a := n.DataAtom
	if n.Type == html.TextNode || isPhrasingElem(n) {
		return true
	}
	isException := (a == atom.A || a == atom.Del || a == atom.Ins)
	for _, child := range children(n) {
		if isPhrasingElem(child) == false {
			return false
		}
	}

	return isException
}

func isPhrasingElem(n *html.Node) bool {
	a := n.DataAtom
	return a == atom.Abbr ||
		a == atom.Audio ||
		a == atom.B ||
		a == atom.Bdo ||
		a == atom.Br ||
		a == atom.Button ||
		a == atom.Cite ||
		a == atom.Code ||
		a == atom.Data ||
		a == atom.Datalist ||
		a == atom.Dfn ||
		a == atom.Em ||
		a == atom.Embed ||
		a == atom.I ||
		a == atom.Img ||
		a == atom.Input ||
		a == atom.Kbd ||
		a == atom.Label ||
		a == atom.Mark ||
		a == atom.Math ||
		a == atom.Meter ||
		a == atom.Noscript ||
		a == atom.Object ||
		a == atom.Output ||
		a == atom.Progress ||
		a == atom.Q ||
		a == atom.Ruby ||
		a == atom.Samp ||
		a == atom.Script ||
		a == atom.Select ||
		a == atom.Small ||
		a == atom.Span ||
		a == atom.Strong ||
		a == atom.Sub ||
		a == atom.Sup ||
		a == atom.Textarea ||
		a == atom.Time ||
		a == atom.Var
}

func children(n *html.Node) []*html.Node {
	matched := make([]*html.Node, 0)
	// Base case
	for f := n.FirstChild; f != nil; f = f.NextSibling {
		matched = append(matched, f)
	}
	return matched
}

func IsScript(n *html.Node) bool {
	a := n.DataAtom
	return a == atom.Script ||
		a == atom.Noscript
}

func IsStyle(n *html.Node) bool {
	a := n.DataAtom
	return a == atom.Style ||
		a == atom.Link
}

func IsBr(n *html.Node) bool {
	return n != nil && n.DataAtom == atom.Br
}

func IsH1(n *html.Node) bool {
	return n.DataAtom == atom.H1
}

func IsP(n *html.Node) bool {
	return n.DataAtom == atom.P
}
