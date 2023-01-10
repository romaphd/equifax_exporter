package InformationPart

import (
	"go_pro/model_jsn"
	"go_pro/model_xml"
	"reflect"
	"strconv"
)

func InformationPart(jsn model_jsn.Model_jsn, event string, blocks []string) *model_xml.InformationPart {
	var xml = model_xml.InformationPart{}
	if len(jsn[0].B55) != 0 && contains(blocks, "55") {
		xml.Application = &model_xml.Application{}
		xml.Application.Ratio = jsn[0].Link[0].CustRoleCode
		xml.Application.Sum = strconv.FormatFloat(jsn[0].B55[0].ReqstCredCcyAmt, 'f', 2, 64)
		xml.Application.Currency = jsn[0].B55[0].ReqstCredCrncyCode
		xml.Application.Uid = jsn[0].Link[0].LiabUniqSid
		xml.Application.Date = jsn[0].B55[0].ReqstDt
		xml.Application.SourceType = jsn[0].B55[0].SrcCode
		xml.Application.Way = jsn[0].B55[0].ReqstCode
		xml.Application.ApprovalDate = jsn[0].B55[0].ReqstAprvdEndDt
	}
	if len(jsn[0].B56) != 0 && contains(blocks, "56") {
		xml.Credit = &model_xml.Credit{}
		xml.Credit.Date = jsn[0].B56[0].FinSubjIssueDt
		xml.Credit.Uid = jsn[0].Link[0].LiabUniqSid
	}

	if len(jsn[0].B57) != 0 && contains(blocks, "57") {
		xml.Failure = &model_xml.Failure{}
		xml.Failure.Date = jsn[0].B57[0].RjctDt
		xml.Failure.Reason = jsn[0].B57[0].ReasonRjctCode
	}

	if reflect.DeepEqual(xml, model_xml.InformationPart{}) {
		return nil
	} else {
		return &xml
	}
}

func contains(a []string, v string) bool {
	for _, i := range a {
		if i == v {
			return true
		}
	}
	return false
}
