type Envelope struct {
	XMLName xml.Name `xml:"soapenv:Envelope"`
	Soapenv string   `xml:"xmlns:soapenv,attr"`
	Java    string   `xml:"xmlns:java,attr"`
	Java1   string   `xml:"xmlns:java1,attr"`
	Body    Body     `xml:"soapenv:Body"`
}

type Body struct {
	XMLName xml.Name `xml:"soapenv:Body"`
	request request  `xml:"java:request"`
}

type request struct {
	XMLName          xml.Name `xml:"java:request"`
	Header           string   `xml:"soapenv:Header"`
	RefnId           string   `xml:"java1:RefnId"`
	subscriberNumber string   `xml:"java1:subscriberNumber"`
	customerNumber   string   `xml:"java1:customerNumber"`
}

