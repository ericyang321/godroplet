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
<head>
	<script type="text/javascript">
		var i = 10;
	</script>
</head>
<body>
	<h1 class="outer-h1" disabled>
		<div>
			<h1 class="inner-h1">
				Tada
			</h1>
		</div>
		<script src="https://code.jquery.com/jquery-3.3.1.min.js" integrity="sha256-FgpCb/KJQlLNfOu91ta32o/NMZxltwRo8QtmkMRdAu8=" crossorigin="anonymous"></script>
	</h1>
</body>
</html>
`

// tests
func TestFindAllShallow(t *testing.T) {
	// setup
	rootNode, parseErr := html.Parse(strings.NewReader(doublyNestedH1))
	if parseErr != nil {
		t.Errorf("HTML Parse error: %s", parseErr.Error())
		return
	}

	// method run
	nodeList := FindAllShallow(rootNode, isH1)

	// result comparison
	for _, nodePointer := range nodeList {
		n := *nodePointer
		if n.DataAtom != atom.H1 {
			t.Errorf("Expected nodes found to all be H1, but instead found a %s", n.Data)
			return
		}
	}
	l := len(nodeList)
	if l != 1 {
		t.Errorf("Expected 1 H1 nodes to be found, but instead found %d", l)
	}
}

func TestFindAllDeep(t *testing.T) {
	// setup
	rootNode, parseErr := html.Parse(strings.NewReader(doublyNestedH1))
	if parseErr != nil {
		t.Errorf("HTML Parse error: %s", parseErr.Error())
		return
	}

	// method run
	nodeList := FindAllDeep(rootNode, isH1)

	// result comparison
	for _, nodePointer := range nodeList {
		n := *nodePointer
		if n.DataAtom != atom.H1 {
			t.Errorf("Expected nodes found to all be H1, but instead found a %s", n.Data)
			return
		}
	}
	l := len(nodeList)
	if l != 2 {
		t.Errorf("Expected 2 H1 nodes to be found, but instead found %d", l)
	}
}

func TestRemoveScriptNodes(t *testing.T) {
	// setup
	rootNode, parseErr := html.Parse(strings.NewReader(doublyNestedH1))
	if parseErr != nil {
		t.Errorf("HTML Parse error: %s", parseErr.Error())
		return
	}

	// method run
	RemoveScriptNodes(rootNode)
	nodeList := FindAllDeep(rootNode, isScriptNode)

	// result comparison
	for _, np := range nodeList {
		if isScriptNode(np) {
			t.Errorf("Expected no script nodes to be found, but instead found some: %v", nodeList)
		}
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
