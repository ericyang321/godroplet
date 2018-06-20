package trim

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ericyang321/godroplet/src/trim/helper"
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
	for _, br := range nodes {
		next := br.NextSibling
		replaced := false
		// delete chained <br>s after the first <br> found
		for ; isBr(next); next = next.NextSibling {
			next = nextElement(next)
			replaced = true
			removeNode(next)
		}
		// if we deleted a chain of <br>s, replace remaining / undeleted <br> with <p>
		if replaced == true {
			p := createElement("<p></p>")
			replace(p, br)

			next = p.NextSibling
			for next != nil {
				// If we've hit another damn <br>, then we're done adding children to this <p>
				if isBr(next) {
					pNextElem := nextElement(next.NextSibling)
					if isBr(pNextElem) {
						break
					}
				}
				if !helper.IsPhrasingContent(next) {
					break
				}
				// Otherwise, make this node a child of the new <p>.
				sibling := next.NextSibling
				removeNode(next)
				p.AppendChild(next)
				next = sibling
			}
		}
	}
}

// ------------------------------------------------------
// --------------------- HELPERS ------------------------
// ------------------------------------------------------

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
	return n != nil && n.DataAtom == atom.Br
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

func createElement(s string) *html.Node {
	return parseHTML(s)[0]
}

func parseHTML(s string) []*html.Node {
	n, err := html.ParseFragment(strings.NewReader(s), &html.Node{
		Type: html.ElementNode,
	})
	if err != nil {
		panic(fmt.Sprintf("Failed to parse HTML: %s", err.Error()))
	}
	return n
}

func replace(newChild, oldChild *html.Node) {
	oldChild.Parent.InsertBefore(newChild, oldChild)
	removeNode(oldChild)
}
