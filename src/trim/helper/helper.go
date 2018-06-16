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

// RemoveNode deletes a selected node from DOM Tree
func RemoveNode(n *html.Node) error {
	parentNode := n.Parent
	if parentNode == nil {
		return fmt.Errorf("Node cannot be deleted without an existing parent node")
	}
	parentNode.RemoveChild(n)
	return nil
}

// RemoveScriptNodes iterates html document and delete all script nodes
func RemoveScriptNodes(n *html.Node) {
	scriptNodes := FindAllShallow(n, isScriptNode)
	for _, np := range scriptNodes {
		RemoveNode(np)
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
