package helper

import (
	"fmt"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Matcher is user passable function to decide which nodes to fetch
type Matcher func(*html.Node) bool

// FindAllShallow finds all nodes that return true from matcher comparison.
// FindAllShallow will not continue searching matching subnodes once node has been matched.
func FindAllShallow(n *html.Node, matcher Matcher) []*html.Node {
	return depthFirstSearch(n, matcher, false)
}

// FindAllDeep will find all nodes that return true from matcher comparison.
// FindAllDeep will continue searching subnodes after node has been matched.
func FindAllDeep(n *html.Node, matcher Matcher) []*html.Node {
	return depthFirstSearch(n, matcher, true)
}

// RemoveScriptStyleNodes iterates DOM tree and deletes all script, style, link, and noscript nodes
func RemoveScriptStyleNodes(n *html.Node) {
	nodes := FindAllShallow(n, isScriptStyleNode)
	for _, np := range nodes {
		removeNode(np)
	}
}

// ------------------------------------------------------
// --------------------- HELPERS ------------------------
// ------------------------------------------------------

// depthFirstSearch recursively finds node pointers and pushes them onto a returned list
func depthFirstSearch(n *html.Node, matcher Matcher, deep bool) []*html.Node {
	matched := make([]*html.Node, 0)
	// Base case
	if matcher(n) {
		matched = append(matched, n)
		if deep == false {
			return matched
		}
	}
	for f := n.FirstChild; f != nil; f = f.NextSibling {
		childNodes := depthFirstSearch(f, matcher, deep)
		if len(childNodes) > 0 {
			matched = append(matched, childNodes...)
		}
	}
	return matched
}

func isScriptNode(n *html.Node) bool {
	a := n.DataAtom
	return a == atom.Script ||
		a == atom.Noscript
}

func isStyleNode(n *html.Node) bool {
	a := n.DataAtom
	return a == atom.Style ||
		a == atom.Link
}

func isScriptStyleNode(n *html.Node) bool {
	return compose(isStyleNode, isScriptNode)(n)
}

func compose(matchers ...Matcher) func(n *html.Node) bool {
	return func(n *html.Node) bool {
		for _, matcher := range matchers {
			if m := matcher(n); m == false {
				return false
			}
		}
		return true
	}
}

func removeNode(n *html.Node) error {
	parentNode := n.Parent
	if parentNode == nil {
		return fmt.Errorf("Node cannot be deleted without an existing parent node")
	}
	parentNode.RemoveChild(n)
	return nil
}
