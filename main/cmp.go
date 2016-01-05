type Commands struct {
    Debug string `xml:"xsd:debug"`
    AudienceID Xsd:AudienceID `xml:"xsd:audienceID"`
    EventParameters Xsd:EventParameters `xml:"xsd:eventParameters"`
    RelyOnExistingSession string `xml:"xsd:relyOnExistingSession"`
    AudienceLevel string `xml:"xsd:audienceLevel"`
    InteractiveChannel string `xml:"xsd:interactiveChannel"`
    MethodIdentifier string `xml:"xsd:methodIdentifier"`
}

type EventParameters struct {
    Name string `xml:"xsd:name"`
    ValueAsString string `xml:"xsd:valueAsString"`
    ValueDataType string `xml:"xsd:valueDataType"`
    ValueAsNumeric string `xml:"xsd:valueAsNumeric"`
}

type ExecuteBatch struct {
    Commands Soap1:Commands `xml:"soap1:commands"`
    SessionID string `xml:"soap1:sessionID"`
}

type Envelope struct {
    Header string `xml:"soap:Header"`
    Body Soap:Body `xml:"soap:Body"`
    Soap string `xml:"xmlns:soap,attr"`
    Soap1 string `xml:"xmlns:soap1,attr"`
    Xsd string `xml:"xmlns:xsd,attr"`
}

type AudienceID struct {
    ValueDataType string `xml:"xsd:valueDataType"`
    Name string `xml:"xsd:name"`
    ValueAsString string `xml:"xsd:valueAsString"`
}

type Body struct {
    ExecuteBatch Soap1:ExecuteBatch `xml:"soap1:executeBatch"`
}

