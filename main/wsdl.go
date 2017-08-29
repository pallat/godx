package main

type Definitions struct {
	Binding         Binding  `xml:"binding"`
	Types           Types    `xml:"types"`
	Service         Service  `xml:"service"`
	Name            string   `xml:"name,attr"`
	Tns             string   `xml:"tns,attr"`
	Xsd1            string   `xml:"xsd1,attr"`
	Xmlns           string   `xml:"xmlns,attr"`
	Message         Message  `xml:"message"`
	PortType        PortType `xml:"portType"`
	TargetNamespace string   `xml:"targetNamespace,attr"`
	Soap            string   `xml:"soap,attr"`
}
type Binding struct {
	Type      string    `xml:"type,attr"`
	Style     string    `xml:"style,attr"`
	Transport string    `xml:"transport,attr"`
	Operation Operation `xml:"operation"`
	Binding   string    `xml:"binding"`
	Name      string    `xml:"name,attr"`
}
type Operation struct {
	Input      Input  `xml:"input"`
	Output     Output `xml:"output"`
	Operation  string `xml:"operation"`
	Name       string `xml:"name,attr"`
	SoapAction string `xml:"soapAction,attr"`
}
type Input struct {
	Body    string `xml:"body"`
	Message string `xml:"message,attr"`
}
type Output struct {
	Body    string `xml:"body"`
	Message string `xml:"message,attr"`
}
type Types struct {
	Schema Schema `xml:"schema"`
}
type Schema struct {
	Element         Element `xml:"element"`
	TargetNamespace string  `xml:"targetNamespace,attr"`
	Xmlns           string  `xml:"xmlns,attr"`
}
type Element struct {
	Name        string      `xml:"name,attr"`
	Type        string      `xml:"type,attr"`
	ComplexType ComplexType `xml:"complexType"`
}
type ComplexType struct {
	All All `xml:"all"`
}
type All struct {
	Element Element `xml:"element"`
}
type Service struct {
	Documentation string `xml:"documentation"`
	Name          string `xml:"name,attr"`
	Port          Port   `xml:"port"`
}
type Port struct {
	Address string `xml:"address"`
	Name    string `xml:"name,attr"`
	Binding string `xml:"binding,attr"`
}
type Message struct {
	Part string `xml:"part"`
	Name string `xml:"name,attr"`
}
type PortType struct {
	Operation Operation `xml:"operation"`
	Name      string    `xml:"name,attr"`
}
