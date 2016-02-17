package godx

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"io/ioutil"
	"testing"
)

func TestAstFunc(t *testing.T) {
	b, err := ioutil.ReadFile("./main/cmp.xml")
	if err != nil {
		t.Error("File cmp.xml not found.")
		return
	}

	x2go := New(b)

	fmt.Println(x2go.MakePackage("pallat"))
}

func TestAstPackage(t *testing.T) {
	fset := token.NewFileSet()

	decls := []ast.Decl{}

	decl := &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{
			makeStruct("Envelope", []*ast.Field{
				makeField("Header", "string", "`xml:\"xsd:Header\"`"),
				makeField("Body", "Body", "`xml:\"xsd:Body\"`"),
				makeField("Soap", "string", "`xml:\"xmlns:soap,attr\"`"),
				makeField("Soap1", "string", "`xml:\"xmlns:soap1,attr\"`"),
				makeField("Xsd", "string", "`xml:\"xmlns:xsd,attr\"`")}),
		},
	}
	decls = append(decls, decl)

	decl = &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{
			makeStruct("Body", []*ast.Field{
				makeField("ExecuteBatch", "ExecuteBatch", "`xml:\"soap1:executeBatch\"`")}),
		},
	}
	decls = append(decls, decl)

	decl = &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{
			makeStruct("ExecuteBatch", []*ast.Field{
				makeField("SessionID", "string", "`xml:\"soap1:sessionID\"`"),
				makeField("Commands", "Commands", "`xml:\"soap1:commands\"`")}),
		},
	}
	decls = append(decls, decl)

	decl = &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{
			makeStruct("Commands", []*ast.Field{
				makeField("EventParameters", "EventParameters", "`xml:\"xsd:eventParameters\"`"),
				makeField("MethodIdentifier", "string", "`xml:\"xsd:methodIdentifier\"`"),
				makeField("AudienceID", "AudienceID", "`xml:\"xsd:audienceID\"`"),
				makeField("AudienceLevel", "string", "`xml:\"xsd:audienceLevel\"`"),
				makeField("InteractiveChannel", "string", "`xml:\"xsd:interactiveChannel\"`"),
				makeField("RelyOnExistingSession", "string", "`xml:\"xsd:relyOnExistingSession\"`"),
				makeField("Debug", "string", "`xml:\"xsd:debug\"`")}),
		},
	}
	decls = append(decls, decl)

	decl = &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{
			makeStruct("EventParameters", []*ast.Field{
				makeField("Name", "string", "`xml:\"xsd:name\"`"),
				makeField("ValueDataType", "string", "`xml:\"xsd:valueDataType\"`"),
				makeField("ValueAsString", "string", "`xml:\"xsd:valueAsString\"`"),
				makeField("ValueAsNumeric", "string", "`xml:\"xsd:valueAsNumeric\"`")}),
		},
	}
	decls = append(decls, decl)

	decl = &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{
			makeStruct("AudienceID", []*ast.Field{
				makeField("Name", "string", "`xml:\"xsd:name\"`"),
				makeField("ValueDataType", "string", "`xml:\"xsd:valueDataType\"`"),
				makeField("ValueAsString", "string", "`xml:\"xsd:valueAsString\"`")}),
		},
	}
	decls = append(decls, decl)

	f := &ast.File{
		Name: &ast.Ident{
			NamePos: 1,
			Name:    "xxx",
		},
		Decls: decls,
	}

	ast.Print(fset, f)

	var buf bytes.Buffer
	printer.Fprint(&buf, fset, f)
	s := buf.String()
	fmt.Println(s)
}
