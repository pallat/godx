package main

type Body struct {
	InqrDiscAmnt InqrDiscAmnt `xml:"Body"`
}
type InqrDiscAmntReqs struct {
	ReqsDttm	string	`xml:"InqrDiscAmntReqs"`
	CmpgId		string	`xml:"InqrDiscAmntReqs"`
	CustNumb	string	`xml:"InqrDiscAmntReqs"`
	Field01		string	`xml:"InqrDiscAmntReqs"`
	RefnId		string	`xml:"InqrDiscAmntReqs"`
	Cmnd		string	`xml:"InqrDiscAmntReqs"`
	Field02		string	`xml:"InqrDiscAmntReqs"`
}
type InqrDiscAmnt struct {
	InqrDiscAmntReqs InqrDiscAmntReqs `xml:"InqrDiscAmnt"`
}
type Envelope struct {
	Body	Body	`xml:"Envelope"`
	Header	string	`xml:"Envelope"`
	Soapenv	string	`xml:"Envelope"`
	Ws	string	`xml:"Envelope"`
}
