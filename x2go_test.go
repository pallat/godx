package x2go

import "testing"

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

var xstruct1 = `type authen struct {
	XMLName xml.Name ` + "`" + `xml:"tem:Authen"` + "`" + `
	Username string ` + "`" + `xml:"tem:Username"` + "`" + `
	Password string ` + "`" + `xml:"tem:Password"` + "`" + `
	DomainName string ` + "`" + `xml:"tem:DomainName"` + "`" + `
	ClientIP string ` + "`" + `xml:"tem:ClientIP"` + "`" + `
}

type body struct {
	XMLName xml.Name ` + "`" + `xml:"soapenv:Body"` + "`" + `
	Authen Authen ` + "`" + `xml:"tem:Authen"` + "`" + `
}

type envelope struct {
	XMLName xml.Name ` + "`" + `xml:"soapenv:Envelope"` + "`" + `
	Soapenv string ` + "`" + `xml:"xmlns:soapenv,attr"` + "`" + `
	Tem string ` + "`" + `xml:"xmlns:tem,attr"` + "`" + `
	Body body ` + "`" + `xml:"soapenv:Body"` + "`" + `
}`
