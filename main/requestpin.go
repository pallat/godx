package main

type Body struct {
	InqrDiscAmnt InqrDiscAmnt `xml:"Body"`
}
type InqrDiscAmnt struct {
	InqrDiscAmntReqs InqrDiscAmntReqs `xml:"InqrDiscAmnt"`
}
type InqrDiscAmntReqs struct {
	Field02		string	`xml:"InqrDiscAmntReqs"`
	CmpgId		string	`xml:"InqrDiscAmntReqs"`
	Cmnd		string	`xml:"InqrDiscAmntReqs"`
	CustNumb	string	`xml:"InqrDiscAmntReqs"`
	Field01		string	`xml:"InqrDiscAmntReqs"`
	ReqsDttm	string	`xml:"InqrDiscAmntReqs"`
	RefnId		string	`xml:"InqrDiscAmntReqs"`
}
type Envelope struct {
	Body		Body	`xml:"Envelope"`
	Header		string	`xml:"Envelope"`
	Soapenv,Attr	string	`xml:"Envelope"`
	Ws,Attr		string	`xml:"Envelope"`
}
