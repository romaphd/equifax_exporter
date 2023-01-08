package model_xml

import (
	"encoding/xml"
)

type InformationPart struct {
	XMLName     xml.Name `xml:"information_part,omitempty"`
	Application *Application
	Failure     *Failure
	Credit      *Credit
}
type Application struct {
	XMLName      xml.Name `xml:"application,omitempty"`
	Ratio        string   `xml:"ratio,omitempty"`
	Sum          string   `xml:"sum,omitempty"`
	Currency     string   `xml:"currency,omitempty"`
	Uid          string   `xml:"uid,omitempty"`
	Date         string   `xml:"date,omitempty"`
	SourceType   string   `xml:"source_type,omitempty"`
	Way          string   `xml:"way,omitempty"`
	ApprovalDate string   `xml:"approval_date,omitempty"`
}

type Credit struct {
	XMLName      xml.Name `xml:"credit,omitempty"`
	Ratio        string   `xml:"ratio,omitempty"`
	Type         string   `xml:"type,omitempty"`
	Uid          string   `xml:"uid,omitempty"`
	Date         string   `xml:"date,omitempty"`
	Sign90plus   string   `xml:"sign_90plus,omitempty"`
	SignStopLoad string   `xml:"sign_stop_load,omitempty"`
}
type Failure struct {
	XMLName xml.Name `xml:"failure,omitempty"`
	Date    string   `xml:"date,omitempty"`
	Reason  string   `xml:"reason,omitempty"`
}
