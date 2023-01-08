package model_xml

import (
	"encoding/xml"
)

type Fch struct {
	XMLName                   xml.Name `xml:"fch"`
	Xsi                       string   `xml:"xsi,attr"`
	NoNamespaceSchemaLocation string   `xml:"noNamespaceSchemaLocation,attr"`
	Version                   string   `xml:"version,attr"`
	Head                      struct {
		SourceInn   string `xml:"source_inn"`
		SourceOgrn  string `xml:"source_ogrn"`
		Date        string `xml:"date"`
		FileRegDate string `xml:"file_reg_date"`
		FileRegNum  string `xml:"file_reg_num"`
		PrevFile    *PrevFile
	} `xml:"head"`
	Info   []Info
	Footer struct {
		SubjectsCount string `xml:"subjects_count"`
		RecordsCount  string `xml:"records_count"`
	} `xml:"footer"`
}

type PrevFile struct {
	XMLName     xml.Name `xml:"prev_file,omitempty"`
	FileRegDate string   `xml:"file_reg_date"`
	FileRegNum  string   `xml:"file_reg_num"`
}

type Info struct {
	XMLName         xml.Name `xml:"info,omitempty"`
	Recnumber       string   `xml:"recnumber,attr,omitempty"`
	Event           string   `xml:"event,attr,omitempty"`
	Action          string   `xml:"action,attr,omitempty"`
	ActionReason    string   `xml:"action_reason,attr,omitempty"`
	PrevFileRegDate string   `xml:"prev_file_reg_date,attr,omitempty"`
	PrevFileRegNum  string   `xml:"prev_file_reg_num,attr,omitempty"`
	Comment         string   `xml:"comment,attr,omitempty"`
	TitlePart       TitlePart
	BasePart        *BasePart
	AddPart         *AddPart
	InformationPart *InformationPart
}

func Contains(a []string, v string) bool {
	for _, i := range a {
		if i == v {
			return true
		}
	}
	return false
}
