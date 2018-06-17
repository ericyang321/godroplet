package helper

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

var exampleHTML = `
<!DOCTYPE html>
<html>
<head>
	<script type="text/javascript">
		var i = 10;
	</script>
	<style>
		h1 {color:red;}
		p {color:blue;}
	</style>
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

var rootNode, _ = html.Parse(strings.NewReader(exampleHTML))

// tests
func TestFindAllShallow(t *testing.T) {
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

func TestRemoveScriptStyleNodes(t *testing.T) {
	// method run
	RemoveScriptStyleNodes(rootNode)
	nodeList := FindAllDeep(rootNode, isScriptStyleNode)

	// result comparison
	for _, np := range nodeList {
		if isScriptStyleNode(np) {
			t.Errorf("Expected no script nodes to be found, but instead found some: %v", nodeList)
		}
	}
}

// ------------------------------------------------------
// --------------------- HELPERS ------------------------
// ------------------------------------------------------

func isH1(n *html.Node) bool {
	return n.DataAtom == atom.H1
}
