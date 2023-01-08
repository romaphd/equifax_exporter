package model_xml

import (
	"encoding/xml"
)

type AddPart struct {
	XMLName    xml.Name `xml:"add_part,omitempty"`
	Accounting struct {
		XMLName xml.Name `xml:"accounting,omitempty"`
		Sign    string   `xml:"sign"`
	}
}
