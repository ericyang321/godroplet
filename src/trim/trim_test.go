package trim

import (
	"fmt"
	"strings"
	"testing"

	h "github.com/ericyang321/godroplet/src/trim/helper"
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
		h1 {
			color:red;
		}
		p {
			color:blue;
		}
	</style>
</head>
<body>
	<h1 class="outer-h1" disabled>
		words
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

var rootNode = getRoot(exampleHTML)

// tests
func TestFindAllShallow(t *testing.T) {
	// method run
	nodeList := FindAllShallow(rootNode, h.IsH1)

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
	nodeList := FindAllDeep(rootNode, h.IsH1)

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
	isScriptAndStyle := Compose(h.IsScript, h.IsStyle)
	// method run
	RemoveScriptStyleNodes(rootNode)
	nodeList := FindAllDeep(rootNode, isScriptAndStyle)

	// result comparison
	for _, np := range nodeList {
		if isScriptAndStyle(np) {
			t.Errorf("Expected no script nodes to be found, but instead found some: %v", nodeList)
		}
	}
}

func TestReplaceBrs(t *testing.T) {
	// method run
	chainedBrs := `<div>foo<br>bar <br><br><br>abc</div>`
	root := getRoot(chainedBrs)

	// result comparison
	// EXPECTED: <div>foo<br>bar<p>abc</p></div>
	ReplaceBrs(root)

	brNodes := FindAllDeep(root, h.IsBr)
	lenBr := len(brNodes)
	if lenBr != 1 {
		t.Errorf("Expected 1 <br> present, but instead found %d", lenBr)
	}
	pNodes := FindAllDeep(root, h.IsP)
	lenP := len(pNodes)
	if lenP != 1 {
		t.Errorf("Expected 1 <p> present, but instead found %d", lenP)
	}
}

// ------------------------------------------------------
// --------------------- HELPERS ------------------------
// ------------------------------------------------------

func getRoot(str string) *html.Node {
	root, err := html.Parse(strings.NewReader(str))
	if err != nil {
		panic(fmt.Sprintf("HTML Parsing error: %s", err.Error()))
	}
	return root
}
