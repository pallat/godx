package godx

import (
	"io/ioutil"
	"testing"
)

func TestAst(t *testing.T) {
	b, err := ioutil.ReadFile("./main/cmp.xml")
	if err != nil {
		t.Error("File cmp.xml not found.")
		return
	}

	x2go := New(b)

	x2go.ast()
}
