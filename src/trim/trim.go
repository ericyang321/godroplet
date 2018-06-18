package trim

import (
	"fmt"
	"regexp"

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

// RemoveScriptStyleNodes iterates DOM tree and deletes all <script>, <style>, <link>, and <noscript> nodes
func RemoveScriptStyleNodes(n *html.Node) {
	nodes := FindAllShallow(n, isScriptStyle)
	for _, np := range nodes {
		removeNode(np)
	}
}

// ReplaceBrs replaces 2 or more consecutive <br> nodes with a single <p> node
// whitespace between <br> elements are ignored.
func ReplaceBrs(n *html.Node) {
	nodes := FindAllDeep(n, isBr)
	for _, np := range nodes {
		next := np.NextSibling
		replaced := false
		// delete chained <br>s after the first <br> found
		for next = nextElement(next); next != nil && next.DataAtom == atom.Br; next = next.NextSibling {
			replaced = true
			removeNode(next)
		}
		// if we deleted a chain of <br>s, replace remaining / undeleted <br>s with <p>
		if replaced == true {
		}
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

func isScript(n *html.Node) bool {
	a := n.DataAtom
	return a == atom.Script ||
		a == atom.Noscript
}

func isStyle(n *html.Node) bool {
	a := n.DataAtom
	return a == atom.Style ||
		a == atom.Link
}

func isBr(n *html.Node) bool {
	return n.DataAtom == atom.Br
}

func isScriptStyle(n *html.Node) bool {
	return compose(isStyle, isScript)(n)
}

func compose(matchers ...Matcher) Matcher {
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

func attr(n *html.Node, key string) string {
	for _, a := range n.Attr {
		if a.Key == key {
			return a.Val
		}
	}
	return ""
}

// Finds the next element, starting from the given node, and ignoring
// whitespace in between. If the given node is an element, the same node is
// returned.
func nextElement(n *html.Node) *html.Node {
	next := n
	for {
		onlySpaces := false
		if next != nil {
			onlySpaces, _ = regexp.MatchString(`^\s*$`, next.Data)
		}
		if next != nil && next.Type != html.ElementNode && onlySpaces == true {
			next = next.NextSibling
		} else {
			break
		}
	}
	return next
}
