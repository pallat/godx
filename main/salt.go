type Envelope struct {
	XMLName xml.Name `xml:"soapenv:Envelope"`
	Soapenv string   `xml:"xmlns:soapenv,attr"`
	Urn     string   `xml:"xmlns:urn,attr"`
	Body    Body     `xml:"soapenv:Body"`
}

type Body struct {
	XMLName           xml.Name          `xml:"soapenv:Body"`
	ReadCustSubInfoWS ReadCustSubInfoWS `xml:":ReadCustSubInfoWS"`
}

type ReadCustSubInfoWS struct {
	XMLName xml.Name `xml:":ReadCustSubInfoWS"`
	inbuf   inbuf    `xml:"urn:inbuf"`
}

type inbuf struct {
	XMLName   xml.Name `xml:"urn:inbuf"`
	Header    string   `xml:"soapenv:Header"`
	USER_CODE string   `xml:":USER_CODE"`
	CUST_NUMB string   `xml:":CUST_NUMB"`
	SUBR_NUMB string   `xml:":SUBR_NUMB"`
}

