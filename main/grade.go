type Envelope struct {
	XMLName xml.Name `xml:"soapenv:Envelope"`
	Soapenv string   `xml:"xmlns:soapenv,attr"`
	Wsd     string   `xml:"xmlns:wsd,attr"`
	Body    Body     `xml:"soapenv:Body"`
}

type Body struct {
	XMLName                       xml.Name                      `xml:"soapenv:Body"`
	QueryCustomerPrivilegeRequest QueryCustomerPrivilegeRequest `xml:"wsd:QueryCustomerPrivilegeRequest"`
}

type QueryCustomerPrivilegeRequest struct {
	XMLName          xml.Name   `xml:"wsd:QueryCustomerPrivilegeRequest"`
	HeaderData       HeaderData `xml:"wsd:HeaderData"`
	SubscriberNumber string     `xml:":SubscriberNumber"`
	CustomerNumber   string     `xml:":CustomerNumber"`
}

type HeaderData struct {
	XMLName       xml.Name `xml:"wsd:HeaderData"`
	Header        string   `xml:"soapenv:Header"`
	MessageId     string   `xml:":MessageId"`
	BusinessEvent string   `xml:":BusinessEvent"`
	SentDateTime  string   `xml:":SentDateTime"`
	SourceSystem  string   `xml:":SourceSystem"`
	ResponseTime  string   `xml:":ResponseTime"`
}

