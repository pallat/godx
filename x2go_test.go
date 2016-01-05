package x2go

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestPrintStruct(t *testing.T) {
	x2go := New([]byte(xexam1))
	s := x2go.String()

	if s != xstruct1 {
		t.Error("expect\n", xstruct1, "but got\n", s)
	}
}

var xexam1 = `<?xml version="1.0"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tem="http://tempuri.org/">
<soapenv:Body>
    <tem:Authen>
        <tem:Username>?</tem:Username>
        <tem:Password>?</tem:Password>
        <tem:DomainName>?</tem:DomainName>
        <tem:ClientIP>?</tem:ClientIP>
    </tem:Authen>
</soapenv:Body>
</soapenv:Envelope>`

var xstruct1 = `type Envelope struct {
	XMLName xml.Name ` + "`" + `xml:"soapenv:Envelope"` + "`" + `
	Soapenv string ` + "`" + `xml:"xmlns:soapenv,attr"` + "`" + `
	Tem string ` + "`" + `xml:"xmlns:tem,attr"` + "`" + `
	Body Body ` + "`" + `xml:"soapenv:Body"` + "`" + `
}

type Body struct {
	XMLName xml.Name ` + "`" + `xml:"soapenv:Body"` + "`" + `
	Authen Authen ` + "`" + `xml:"tem:Authen"` + "`" + `
}

type Authen struct {
	XMLName xml.Name ` + "`" + `xml:"tem:Authen"` + "`" + `
	Username string ` + "`" + `xml:"tem:Username"` + "`" + `
	Password string ` + "`" + `xml:"tem:Password"` + "`" + `
	DomainName string ` + "`" + `xml:"tem:DomainName"` + "`" + `
	ClientIP string ` + "`" + `xml:"tem:ClientIP"` + "`" + `
}

`

func TestLayer(t *testing.T) {
	b, err := ioutil.ReadFile("./main/cmp.xml")
	if err != nil {
		t.Error("File cmp.xml not found.")
		return
	}

	x2go := New(b)

	l := x2go.Layer()
	if l != 5 {
		t.Error("It should return 5 but was ", l)
	}
}

func TestSkeleton(t *testing.T) {
	b, err := ioutil.ReadFile("./main/cmp.xml")
	if err != nil {
		t.Error("File cmp.xml not found.")
		return
	}

	x2go := New(b)

	bones := x2go.Skeleton()

	expect := map[string][]string{
		"":                    []string{"soap:Envelope"},
		"soap:Envelope":       []string{"soap:Header", "soap:Body", "xmlns:soap,attr", "xmlns:soap1,attr", "xmlns:xsd,attr"},
		"soap:Body":           []string{"soap1:executeBatch"},
		"soap1:executeBatch":  []string{"soap1:sessionID", "soap1:commands"},
		"soap1:commands":      []string{"xsd:audienceID", "xsd:audienceLevel", "xsd:debug", "xsd:eventParameters", "xsd:interactiveChannel", "xsd:methodIdentifier", "xsd:relyOnExistingSession"},
		"xsd:eventParameters": []string{"xsd:name", "xsd:valueAsString", "xsd:valueDataType", "xsd:valueAsNumeric"},
	}

	for k, v := range bones.(map[string][]string) {
		if len(v) != len(expect[k]) {
			t.Error("Something went wrong.")
			t.Errorf("%# v:%# v\n!=\n%# v:%# v", k, v, k, expect[k])
		}
	}
	// arrange("", bones.(map[string][]string))
}

func TestIdentifyType(t *testing.T) {
	bones := map[string][]string{
		"":                    []string{"soap:Envelope"},
		"soap:Envelope":       []string{"soap:Header", "soap:Body", "xmlns:soap,attr", "xmlns:soap1,attr", "xmlns:xsd,attr"},
		"soap:Body":           []string{"soap1:executeBatch"},
		"soap1:executeBatch":  []string{"soap1:sessionID", "soap1:commands"},
		"soap1:commands":      []string{"xsd:audienceID", "xsd:audienceLevel", "xsd:debug", "xsd:eventParameters", "xsd:interactiveChannel", "xsd:methodIdentifier", "xsd:relyOnExistingSession"},
		"xsd:eventParameters": []string{"xsd:name", "xsd:valueAsString", "xsd:valueDataType", "xsd:valueAsNumeric"},
	}

	id := Identify(bones)

	expect := map[string]map[string]string{
		"":                map[string]string{"soap:Envelope": "Envelope"},
		"Envelope":        map[string]string{"soap:Header": "Header", "soap:Body": "Body", "xmlns:soap,attr": "string", "xmlns:soap1,attr": "string", "xmlns:xsd,attr": "string"},
		"Body":            map[string]string{"soap1:executeBatch": "ExecuteBatch"},
		"executeBatch":    map[string]string{"soap1:sessionID": "string", "soap1:commands": "Commands"},
		"commands":        map[string]string{"xsd:audienceID": "string", "xsd:audienceLevel": "string", "xsd:debug": "string", "xsd:eventParameters": "EventParameters", "xsd:interactiveChannel": "string", "xsd:methodIdentifier": "string", "xsd:relyOnExistingSession": "string"},
		"eventParameters": map[string]string{"xsd:name": "string", "xsd:valueAsString": "string", "xsd:valueDataType": "string", "xsd:valueAsNumeric": "string"},
	}

	for k, v := range expect {
		if len(v) != len(id[k]) {
			t.Error("It might be wrong.")
			t.Errorf("%# v:%# v\n!=\n%# v:%# v", k, v, k, id[k])
		}
	}
}

func TestPrint(t *testing.T) {
	id := map[string]map[string]string{
		"":                map[string]string{"soap:Envelope": "Envelope"},
		"Envelope":        map[string]string{"soap:Header": "Header", "soap:Body": "Body", "xmlns:soap,attr": "string", "xmlns:soap1,attr": "string", "xmlns:xsd,attr": "string"},
		"Body":            map[string]string{"soap1:executeBatch": "ExecuteBatch"},
		"executeBatch":    map[string]string{"soap1:sessionID": "string", "soap1:commands": "Commands"},
		"commands":        map[string]string{"xsd:audienceID": "string", "xsd:audienceLevel": "string", "xsd:debug": "string", "xsd:eventParameters": "EventParameters", "xsd:interactiveChannel": "string", "xsd:methodIdentifier": "string", "xsd:relyOnExistingSession": "string"},
		"eventParameters": map[string]string{"xsd:name": "string", "xsd:valueAsString": "string", "xsd:valueDataType": "string", "xsd:valueAsNumeric": "string"},
	}

	echo(id)
}

func echo(id map[string]map[string]string) {
	var names []string
	for k, v := range id {
		if k == "" {
			continue
		}

		fmt.Println("type", strings.Title(k), "struct {")
		for name, typ := range v {
			if strings.Contains(name, ",") {
				names = strings.Split(name, ",")
				// names = append(names, names[0])
				names = strings.Split(names[0], ":")
			} else if strings.Contains(name, ":") {
				names = strings.Split(name, ":")
			}
			fmt.Println(" ", strings.Title(names[len(names)-1]), typ, "`xml:"+`"`+name+`"`+"`")
		}
		fmt.Println("}")
	}
}

func arrange(key string, bones map[string][]string) {
	for i := range bones[key] {
		fmt.Println(key + ":" + bones[key][i])
		arrange(bones[key][i], bones)
	}
}
