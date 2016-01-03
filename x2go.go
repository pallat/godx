package x2go

import (
	"bytes"
	"encoding/xml"
	"strings"
)

type X2Go struct {
	dec *xml.Decoder
}

func New(b []byte) *X2Go {
	return &X2Go{
		dec: xml.NewDecoder(bytes.NewReader(b)),
	}
}

type node struct {
	name  xml.Name
	_type string
}

func (x *X2Go) String() string {
	xmlNames := []xml.Name{}
	childs := []node{}
	attrs := map[string]string{}
	start := map[string][]xml.Attr{}
	structs := []string{}

	_struct := ""

	for token, err := x.dec.Token(); err == nil; token, err = x.dec.Token() {
		switch t := token.(type) {
		case xml.StartElement:
			xmlNames = append(xmlNames, t.Name)
			if len(t.Attr) != 0 {
				for _, v := range t.Attr {
					attrs[v.Value] = v.Name.Local
				}
				start[t.Name.Local] = t.Attr
			}
		case xml.EndElement:
			if xmlNames[len(xmlNames)-1] == t.Name {
				childs = append(childs, node{name: t.Name, _type: "string"})
			} else {
				_struct += "type " + t.Name.Local + " struct {\n"
				_struct += "\tXMLName xml.Name `xml:" + `"` + attrs[t.Name.Space] + ":" + t.Name.Local + `"` + "`\n"
				for _, v := range start[t.Name.Local] {
					_struct += "\t" + strings.ToUpper(string(v.Name.Local[0])) + string(v.Name.Local[1:]) + " string" + " `xml:" + `"` + v.Name.Space + ":" + v.Name.Local + `,attr"` + "`\n"
				}

				for _, v := range childs {
					_struct += "\t" + v.name.Local + " " + v._type + " `xml:" + `"` + attrs[v.name.Space] + ":" + v.name.Local + `"` + "`\n"
				}

				_struct += "}\n\n"

				structs = append(structs, _struct)
				_struct = ""

				childs = []node{node{name: t.Name, _type: t.Name.Local}}
			}
		}
	}

	for i := len(structs) - 1; i >= 0; i-- {
		_struct += structs[i]
	}

	return _struct
}

func (x *X2Go) Layer() int {
	count := 0
	names := []string{}

	for token, err := x.dec.Token(); err == nil; token, err = x.dec.Token() {
		switch t := token.(type) {
		case xml.StartElement:
			names = append(names, t.Name.Local)
			if count < len(names) {
				count = len(names)
			}
		case xml.EndElement:
			if t.Name.Local == names[len(names)-1] {
				names = names[:len(names)-1]
			}
		}
	}

	return count - 1
}

func (x *X2Go) Skeleton() interface{} {
	count := 0
	names := []string{}
	mapping := map[string]string{}

	var val = func(s []string) string {
		if len(s) > 1 {
			return s[len(s)-2]
		}
		return ""
	}

	for token, err := x.dec.Token(); err == nil; token, err = x.dec.Token() {
		switch t := token.(type) {
		case xml.StartElement:
			names = append(names, t.Name.Local)
			mapping[names[len(names)-1]] = val(names)
			if count < len(names) {
				count = len(names)
			}
		case xml.EndElement:
			if t.Name.Local == names[len(names)-1] {
				names = names[:len(names)-1]
			}
		}
	}

	bones := map[string][]string{}

	for k, v := range mapping {
		bones[v] = append(bones[v], k)
	}

	return bones
}
