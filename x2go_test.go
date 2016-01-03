package x2go

import (
	"fmt"
	"io/ioutil"
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
		"":                []string{"Envelope"},
		"Envelope":        []string{"Header", "Body"},
		"Body":            []string{"executeBatch"},
		"executeBatch":    []string{"sessionID", "commands"},
		"commands":        []string{"audienceID", "audienceLevel", "debug", "eventParameters", "interactiveChannel", "methodIdentifier", "relyOnExistingSession"},
		"eventParameters": []string{"name", "valueAsString", "valueDataType", "valueAsNumeric"},
	}

	for k, v := range bones.(map[string][]string) {
		if len(v) != len(expect[k]) {
			t.Error("Something went wrong.")
		}
	}
	// arrange("", bones.(map[string][]string))
}

// func TestIdentifyType(t *testing.T) {
// 	bones := map[string][]string{
// 		"":                []string{"Envelope"},
// 		"Envelope":        []string{"Header", "Body"},
// 		"Body":            []string{"executeBatch"},
// 		"executeBatch":    []string{"sessionID", "commands"},
// 		"commands":        []string{"audienceID", "audienceLevel", "debug", "eventParameters", "interactiveChannel", "methodIdentifier", "relyOnExistingSession"},
// 		"eventParameters": []string{"name", "valueAsString", "valueDataType", "valueAsNumeric"},
// 	}
//
// 	id := Identify(bones)
//
// 	expect := []Type{
// 		Type{Name: "Envelope", Type: "Envelope"},
// 		Type{Name: "Body", Type: "Body"},
// 		Type{Name: "Header", Type: "Header"},
// 		Type{Name: "ExecuteBatch", Type: "ExecuteBatch"},
// 		Type{Name: "SessionID", Type: "string"},
// 		Type{Name: "Commands", Type: "Commands"},
// 		Type{Name: "AudienceID", Type: "string"},
// 		Type{Name: "AudienceLevel", Type: "string"},
// 		Type{Name: "Debug", Type: "string"},
// 		Type{Name: "InteractiveChannel", Type: "string"},
// 		Type{Name: "MethodIdentifier", Type: "string"},
// 		Type{Name: "RelyOnExistingSession", Type: "string"},
// 		Type{Name: "EventParameters", Type: "EventParameters"},
// 		Type{Name: "Name", Type: "string"},
// 		Type{Name: "ValueAsString", Type: "string"},
// 		Type{Name: "ValueDataType", Type: "string"},
// 		Type{Name: "ValueAsNumeric", Type: "string"},
// 	}
// }
//
// type Type struct {
// 	Name  string
// 	Type  string
// 	Child []string
// }

func arrange(key string, bones map[string][]string) {
	for i := range bones[key] {
		fmt.Println(key + ":" + bones[key][i])
		arrange(bones[key][i], bones)
	}
}
