package TitlePart

import (
	"fmt"
	"go_pro/model_jsn"
	"go_pro/model_xml"
)

func TitlePart(jsn model_jsn.Model_jsn) model_xml.TitlePart {
	var xml = model_xml.TitlePart{}

	if len(jsn[0].B1) != 0 {
		xml.Private.Name.Last = jsn[0].B1[0].SurnameName
		xml.Private.Name.First = jsn[0].B1[0].FirstName
		xml.Private.Name.Middle = jsn[0].B1[0].FatherName
	}

	if len(jsn[0].B4) != 0 {
		xml.Private.Doc.Country = jsn[0].B4[0].CntryOksmDigitalCode
		xml.Private.Doc.CountryText = jsn[0].B4[0].RegAddrrecretCntryName
		xml.Private.Doc.Type = jsn[0].B4[0].DocKindCode
		xml.Private.Doc.TypeText = jsn[0].B4[0].OtherDocName
		xml.Private.Doc.Serial = jsn[0].B4[0].SerNum
		xml.Private.Doc.Number = jsn[0].B4[0].Num
		xml.Private.Doc.Date = jsn[0].B4[0].IssueDt
		xml.Private.Doc.Who = jsn[0].B4[0].IssueCode
		xml.Private.Doc.DepartmentCode = jsn[0].B4[0].IssueTxt
		xml.Private.Doc.EndDate = jsn[0].B4[0].PlanDt
	}

	if len(jsn[0].B2) != 0 {
		for i := 1; i <= len(jsn[0].B2); i++ {
			xml.Private.History.HistoryName = append(xml.Private.History.HistoryName, model_xml.HistoryName{})
			xml.Private.History.HistoryName[0].First = jsn[0].B2[0].FirstName
		}
	} else {
		xml.Private.History.HistoryName = append(xml.Private.History.HistoryName, model_xml.HistoryName{})
		xml.Private.History.HistoryName[0].Hist_name_sign = "0"
	}

	if len(jsn[0].B5) != 0 {
		for i := 1; i <= len(jsn[0].B5); i++ {
			xml.Private.History.HistoryDoc = append(xml.Private.History.HistoryDoc, model_xml.HistoryDoc{})
			xml.Private.History.HistoryDoc[0].Country = jsn[0].B1[0].SurnameName
		}
	} else {
		xml.Private.History.HistoryDoc = append(xml.Private.History.HistoryDoc, model_xml.HistoryDoc{})
		xml.Private.History.HistoryDoc[0].Hist_doc_sign = "0"
	}

	if len(jsn[0].B6) != 0 {
		xml.Private.Inn.Code = jsn[0].B6[0].InnTypeCode
		xml.Private.Inn.No = jsn[0].B6[0].InnNum
		xml.Private.Inn.Ogrnip = jsn[0].B6[0].RegNum
	}
	if len(jsn[0].B7) != 0 {
		xml.Private.Snils.No = jsn[0].B7[0].SnilsNum
	}

	if len(jsn[0].B3) == 0 {

	} else {
		xml.Private.Birth.Date = DateFormat(jsn[0].B3[0].BirthDt)
	}

	//xml.Private.Birth.Country = jsn[0].B3[0].BirthCntryOksmCode
	//xml.Private.Birth.Place = jsn[0].B3[0].BirthPlaceTxt

	return xml
}

func DateFormat(p_dt string) string {
	if p_dt == "" {
		return "01.01.1900"
	} else {
		return fmt.Sprintf("%v", p_dt)
	}
}
