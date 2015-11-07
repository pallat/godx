package x2go

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"testing"
)

type envelope struct {
	XMLName xml.Name `xml:"soapenv:Envelope"`
	Soapenv string   `xml:"xmlns:soapenv,attr"`
	Tem     string   `xml:"xmlns:tem,attr"`
	Body    body
}

type body struct {
	XMLName xml.Name `xml:"soapenv:Body"`
	Authen  Authen
}

type Authen struct {
	XMLName    xml.Name `xml:"tem:Authen" json:"-"`
	Username   string   `xml:"tem:Username" json:"username"`
	Password   string   `xml:"tem:Password" json:"password"`
	DomainName string   `xml:"tem:DomainName" json:"domain"`
	ClientIP   string   `xml:"tem:-,omitempty" json:"-"`
}

// type authenResponse struct {
// 	AuthenResult string `xml:"Body>AuthenResponse>AuthenResult"`
// }

type node struct {
	name  xml.Name
	_type string
}

func TestPrintStruct(t *testing.T) {
	b, _ := ioutil.ReadFile("req.xml")

	reader := bytes.NewReader(b)

	dec := xml.NewDecoder(reader)

	xmlNames := []xml.Name{}
	childs := []node{}
	attrs := map[string]string{}

	for token, err := dec.Token(); err == nil; token, err = dec.Token() {
		switch t := token.(type) {
		case xml.StartElement:
			xmlNames = append(xmlNames, t.Name)
			if len(t.Attr) != 0 {
				for _, v := range t.Attr {
					attrs[v.Value] = v.Name.Local
				}
			}
		// case xml.CharData:
		// 	// push
		// 	values = append(values, string([]byte(t)))
		case xml.EndElement:
			if xmlNames[len(xmlNames)-1] == t.Name {
				childs = append(childs, node{name: t.Name, _type: "string"})
			} else {
				fmt.Println("struct " + t.Name.Local + " {")
				fmt.Println("\tXMLName xml.Name `xml:" + `"` + attrs[t.Name.Space] + ":" + t.Name.Local + `"` + "`")
				for _, v := range childs {
					fmt.Print("\t" + v.name.Local + " " + v._type + " `xml:" + `"` + attrs[v.name.Space] + ":" + v.name.Local + `"` + "`\n")
				}

				fmt.Println("}")
				childs = []node{node{name: t.Name, _type: t.Name.Local}}
			}

			// m[t.Name.Local] = values[len(values)-1]
			// pop
			// values = values[:len(values)]
		}
	}

}
