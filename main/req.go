type Envelope struct {
	XMLName xml.Name `xml:"soapenv:Envelope"`
	Soapenv string   `xml:"xmlns:soapenv,attr"`
	Tem     string   `xml:"xmlns:tem,attr"`
	Body    Body     `xml:"soapenv:Body"`
}

type Body struct {
	XMLName xml.Name `xml:"soapenv:Body"`
	Authen  Authen   `xml:"tem:Authen"`
}

type Authen struct {
	XMLName    xml.Name `xml:"tem:Authen"`
	Username   string   `xml:"tem:Username"`
	Password   string   `xml:"tem:Password"`
	DomainName string   `xml:"tem:DomainName"`
	ClientIP   string   `xml:"tem:ClientIP"`
}

