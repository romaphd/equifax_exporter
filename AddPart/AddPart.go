package AddPart

import (
	"go_pro/model_jsn"
	"go_pro/model_xml"
	"reflect"
)

func AddPart(jsn model_jsn.Model_jsn, event string, blocks []string) *model_xml.AddPart {
	var xml = model_xml.AddPart{}

	//xml.Accounting.Sign = jsn[0].B8[0].RegAddrrecretFlatNum

	if reflect.DeepEqual(xml, model_xml.AddPart{}) {
		return nil
	} else {
		return &xml
	}
}
