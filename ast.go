package godx

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
	"strings"
)

func makeField(name, _type, tag string) *ast.Field {
	return &ast.Field{
		Names: []*ast.Ident{
			&ast.Ident{
				Name: name,
				Obj:  ast.NewObj(ast.Var, name),
			},
		},
		Type: &ast.Ident{
			Name: _type,
		},
		Tag: &ast.BasicLit{
			Kind:  token.STRING,
			Value: tag,
		},
	}
}

func makeStruct(name string, member []*ast.Field) *ast.TypeSpec {
	return &ast.TypeSpec{
		Name: ast.NewIdent(name),
		Type: &ast.StructType{
			Fields: &ast.FieldList{
				List: member,
			},
		},
	}
}

func makePack(name string, id map[string]map[string]string) string {
	fset := token.NewFileSet()

	decls := []ast.Decl{}

	for k, v := range id {
		if k == "" {
			continue
		}

		decls = append(decls, makeGenDecl(strings.Title(k), v))

	}

	f := &ast.File{
		Name: &ast.Ident{
			NamePos: 1,
			Name:    name,
		},
		Decls: decls,
	}

	var buf bytes.Buffer
	printer.Fprint(&buf, fset, f)
	return buf.String()
}

func makeGenDecl(name string, members map[string]string) *ast.GenDecl {
	fields := []*ast.Field{}
	// var names []string

	for n, typ := range members {
		typs := strings.Split(typ, ":")
		typ = typs[len(typs)-1]

		fields = append(fields, makeField(strings.Title(fieldName(n)), typ, "`xml:\""+name+"\"`"))
	}

	return &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{
			makeStruct(name, fields),
		},
	}
}

func fieldName(n string) string {
	name := n
	if strings.Contains(n, ",") {
		names := strings.Split(n, ",")
		names = strings.Split(names[0], ":")
		name = names[len(names)-1]
	}

	if strings.Contains(n, ":") {
		names := strings.Split(n, ":")
		name = names[len(names)-1]
	}

	return name
}
