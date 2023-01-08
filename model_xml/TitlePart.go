package model_xml

import (
	"encoding/xml"
)

type TitlePart struct {
	XMLName xml.Name `xml:"title_part"`
	Kski    struct {
		Code string `xml:"code,omitempty"`
	} `xml:"kski"`
	Private struct {
		Name struct {
			Last   string `xml:"last,omitempty"`
			First  string `xml:"first,omitempty"`
			Middle string `xml:"middle,omitempty"`
		} `xml:"name"`
		Doc struct {
			Country        string `xml:"country,omitempty"`
			CountryText    string `xml:"country_text,omitempty"`
			Type           string `xml:"type,omitempty"`
			TypeText       string `xml:"type_text,omitempty"`
			Serial         string `xml:"serial,omitempty"`
			Number         string `xml:"number,omitempty"`
			Date           string `xml:"date,omitempty"`
			Who            string `xml:"who,omitempty"`
			DepartmentCode string `xml:"department_code,omitempty"`
			EndDate        string `xml:"end_date,omitempty"`
		} `xml:"doc"`
		Birth struct {
			Date    string `xml:"date"`
			Country string `xml:"country"`
			Place   string `xml:"place"`
		} `xml:"birth"`
		History struct {
			XMLName     xml.Name `xml:"history"`
			HistoryName []HistoryName
			HistoryDoc  []HistoryDoc
		}
		Inn struct {
			Code   string `xml:"code,omitempty"`
			No     string `xml:"no,omitempty"`
			Ogrnip string `xml:"ogrnip,omitempty"`
		} `xml:"inn,omitempty"`
		Snils struct {
			No string `xml:"no,omitempty"`
		} `xml:"snils,omitempty"`
	} `xml:"private,omitempty"`
}

type HistoryName struct {
	XMLName        xml.Name `xml:"name,omitempty"`
	Hist_name_sign string   `xml:"hist_name_sign,omitempty"`
	Last           string   `xml:"last,omitempty"`
	First          string   `xml:"first,omitempty"`
	Middle         string   `xml:"middle,omitempty"`
	DocDate        string   `xml:"doc_date,omitempty"`
}
type HistoryDoc struct {
	XMLName        xml.Name `xml:"doc,omitempty"`
	Hist_doc_sign  string   `xml:"hist_doc_sign,omitempty"`
	Country        string   `xml:"country,omitempty"`
	CountryText    string   `xml:"country_text,omitempty"`
	Type           string   `xml:"type,omitempty"`
	TypeText       string   `xml:"type_text,omitempty"`
	Serial         string   `xml:"serial,omitempty"`
	Number         string   `xml:"number,omitempty"`
	Date           string   `xml:"date,omitempty"`
	Who            string   `xml:"who,omitempty"`
	DepartmentCode string   `xml:"department_code,omitempty"`
	EndDate        string   `xml:"end_date,omitempty"`
}
