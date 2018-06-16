package helper

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

var doublyNestedH1 = `
<!DOCTYPE html>
<html>
<body>
	<h1 class="outer-h1">
		<div>
			<h1 class="inner-h1">
				Tada
			</h1>
		</div>
	</h1>
</body>
</html>
`

// tests
func TestGrep(t *testing.T) {
	// setup
	rootNode, parseErr := html.Parse(strings.NewReader(doublyNestedH1))
	if parseErr != nil {
		t.Errorf("HTML Parse error: %s", parseErr.Error())
		return
	}

	// method run
	nodeList := Grep(rootNode, isH1)
	l := len(nodeList)

	// result comparison
	if l != 1 {
		t.Errorf("Expected 1 H1 nodes to be found, but instead found %d", l)
	}
}

func TestDeepGrep(t *testing.T) {
	// setup
	rootNode, parseErr := html.Parse(strings.NewReader(doublyNestedH1))
	if parseErr != nil {
		t.Errorf("HTML Parse error: %s", parseErr.Error())
		return
	}

	// method run
	nodeList := DeepGrep(rootNode, isH1)
	l := len(nodeList)

	// result comparison
	if l != 2 {
		t.Errorf("Expected 2 H1 nodes to be found, but instead found %d", l)
	}
}

// helpers
func readTestHTML() io.Reader {
	file, openErr := os.Open("./example.html")
	if openErr != nil {
		fmt.Println(openErr.Error())
		os.Exit(2)
	}
	return file
}

func isH1(n *html.Node) bool {
	return n.DataAtom == atom.H1
}
