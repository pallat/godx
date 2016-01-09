type EventParameters struct {
    ValueDataType string `xml:"xsd:valueDataType"`
    ValueAsNumeric string `xml:"xsd:valueAsNumeric"`
    Name string `xml:"xsd:name"`
    ValueAsString string `xml:"xsd:valueAsString"`
}

type AudienceID struct {
    Name string `xml:"xsd:name"`
    ValueAsString string `xml:"xsd:valueAsString"`
    ValueDataType string `xml:"xsd:valueDataType"`
}

type Commands struct {
    AudienceLevel string `xml:"xsd:audienceLevel"`
    Debug string `xml:"xsd:debug"`
    EventParameters EventParameters `xml:"xsd:eventParameters"`
    InteractiveChannel string `xml:"xsd:interactiveChannel"`
    RelyOnExistingSession string `xml:"xsd:relyOnExistingSession"`
    MethodIdentifier string `xml:"xsd:methodIdentifier"`
    AudienceID AudienceID `xml:"xsd:audienceID"`
}

type Envelope struct {
    Header string `xml:"soap:Header"`
    Body Body `xml:"soap:Body"`
    Soap string `xml:"xmlns:soap,attr"`
    Soap1 string `xml:"xmlns:soap1,attr"`
    Xsd string `xml:"xmlns:xsd,attr"`
}

type Body struct {
    ExecuteBatch ExecuteBatch `xml:"soap1:executeBatch"`
}

type ExecuteBatch struct {
    SessionID string `xml:"soap1:sessionID"`
    Commands Commands `xml:"soap1:commands"`
}

