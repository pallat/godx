type Envelope struct {
	XMLName xml.Name `xml:"soap:Envelope"`
	Soap    string   `xml:"xmlns:soap,attr"`
	Ins     string   `xml:"xmlns:ins,attr"`
	Body    Body     `xml:"soap:Body"`
}

type Body struct {
	XMLName        xml.Name       `xml:"soap:Body"`
	DisplayRequest DisplayRequest `xml:"ins:DisplayRequest"`
}

type DisplayRequest struct {
	XMLName           xml.Name          `xml:"ins:DisplayRequest"`
	RequestParameters RequestParameters `xml:"ins:RequestParameters"`
}

type RequestParameters struct {
	XMLName       xml.Name      `xml:"ins:RequestParameters"`
	RequestHeader RequestHeader `xml:"ins:RequestHeader"`
	Parameter     string        `xml:"ins:Parameter"`
	Parameter     string        `xml:"ins:Parameter"`
	Parameter     string        `xml:"ins:Parameter"`
	Parameter     string        `xml:"ins:Parameter"`
	Parameter     string        `xml:"ins:Parameter"`
	Parameter     string        `xml:"ins:Parameter"`
	Parameter     string        `xml:"ins:Parameter"`
}

type RequestHeader struct {
	XMLName xml.Name `xml:"ins:RequestHeader"`
	Header  string   `xml:"soap:Header"`
	NeType  string   `xml:"ins:NeType"`
	ReqUser string   `xml:"ins:ReqUser"`
}

