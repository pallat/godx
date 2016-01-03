type Envelope struct {
	XMLName xml.Name `xml:"S:Envelope"`
	S       string   `xml:"xmlns:S,attr"`
	Body    Body     `xml:"S:Body"`
}

type Body struct {
	XMLName  xml.Name `xml:"S:Body"`
	Response Response `xml:"xmlns:Response"`
}

type Response struct {
	XMLName           xml.Name          `xml:"xmlns:Response"`
	Xmlns             string            `xml:":xmlns,attr"`
	RequestParameters RequestParameters `xml:"xmlns:RequestParameters"`
}

type RequestParameters struct {
	XMLName            xml.Name           `xml:"xmlns:RequestParameters"`
	ResponseParameters ResponseParameters `xml:"xmlns:ResponseParameters"`
	Parameter          string             `xml:"xmlns:Parameter"`
	Parameter          string             `xml:"xmlns:Parameter"`
	Parameter          string             `xml:"xmlns:Parameter"`
	Parameter          string             `xml:"xmlns:Parameter"`
	Parameter          string             `xml:"xmlns:Parameter"`
	Parameter          string             `xml:"xmlns:Parameter"`
	Parameter          string             `xml:"xmlns:Parameter"`
	Parameter          string             `xml:"xmlns:Parameter"`
	Parameter          string             `xml:"xmlns:Parameter"`
	Parameter          string             `xml:"xmlns:Parameter"`
	Parameter          string             `xml:"xmlns:Parameter"`
	Parameter          string             `xml:"xmlns:Parameter"`
	Parameter          string             `xml:"xmlns:Parameter"`
}

type ResponseParameters struct {
	XMLName        xml.Name       `xml:"xmlns:ResponseParameters"`
	ResponseHeader ResponseHeader `xml:"xmlns:ResponseHeader"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
	Parameter      string         `xml:"xmlns:Parameter"`
}

type ResponseHeader struct {
	XMLName         xml.Name `xml:"xmlns:ResponseHeader"`
	RequestId       string   `xml:"xmlns:RequestId"`
	Status          string   `xml:"xmlns:Status"`
	OrderNo         string   `xml:"xmlns:OrderNo"`
	StatusMessage   string   `xml:"xmlns:StatusMessage"`
	StatusMessageId string   `xml:"xmlns:StatusMessageId"`
	Priority        string   `xml:"xmlns:Priority"`
	ReqUser         string   `xml:"xmlns:ReqUser"`
	ReceivedDate    string   `xml:"xmlns:ReceivedDate"`
	FinishedDate    string   `xml:"xmlns:FinishedDate"`
}

