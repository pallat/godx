type Envelope struct {
	XMLName xml.Name `xml:"soap:Envelope"`
	Soap    string   `xml:"xmlns:soap,attr"`
	Soap1   string   `xml:"xmlns:soap1,attr"`
	Xsd     string   `xml:"xmlns:xsd,attr"`
	Body    Body     `xml:"soap:Body"`
}

type Body struct {
	XMLName      xml.Name     `xml:"soap:Body"`
	executeBatch executeBatch `xml:"soap1:executeBatch"`
}

type executeBatch struct {
	XMLName  xml.Name `xml:"soap1:executeBatch"`
	commands commands `xml:"soap1:commands"`
}

type commands struct {
	XMLName          xml.Name `xml:"soap1:commands"`
	commands         commands `xml:"soap1:commands"`
	methodIdentifier string   `xml:"xsd:methodIdentifier"`
}

type commands struct {
	XMLName               xml.Name        `xml:"soap1:commands"`
	eventParameters       eventParameters `xml:"xsd:eventParameters"`
	interactiveChannel    string          `xml:"xsd:interactiveChannel"`
	methodIdentifier      string          `xml:"xsd:methodIdentifier"`
	relyOnExistingSession string          `xml:"xsd:relyOnExistingSession"`
}

type eventParameters struct {
	XMLName         xml.Name        `xml:"xsd:eventParameters"`
	eventParameters eventParameters `xml:"xsd:eventParameters"`
	name            string          `xml:"xsd:name"`
	valueAsNumeric  string          `xml:"xsd:valueAsNumeric"`
	valueDataType   string          `xml:"xsd:valueDataType"`
}

type eventParameters struct {
	XMLName         xml.Name        `xml:"xsd:eventParameters"`
	eventParameters eventParameters `xml:"xsd:eventParameters"`
	name            string          `xml:"xsd:name"`
	valueAsString   string          `xml:"xsd:valueAsString"`
	valueDataType   string          `xml:"xsd:valueDataType"`
}

type eventParameters struct {
	XMLName       xml.Name   `xml:"xsd:eventParameters"`
	audienceID    audienceID `xml:"xsd:audienceID"`
	audienceLevel string     `xml:"xsd:audienceLevel"`
	debug         string     `xml:"xsd:debug"`
	name          string     `xml:"xsd:name"`
	valueAsString string     `xml:"xsd:valueAsString"`
	valueDataType string     `xml:"xsd:valueDataType"`
}

type audienceID struct {
	XMLName       xml.Name `xml:"xsd:audienceID"`
	Header        string   `xml:"soap:Header"`
	sessionID     string   `xml:"soap1:sessionID"`
	name          string   `xml:"xsd:name"`
	valueAsString string   `xml:"xsd:valueAsString"`
	valueDataType string   `xml:"xsd:valueDataType"`
}

