package BasePart

import (
	"go_pro/model_jsn"
	"go_pro/model_xml"
	"reflect"
)

func BasePart(jsn model_jsn.Model_jsn, event string, blocks []string) *model_xml.BasePart {
	var xml = model_xml.BasePart{}

	if len(jsn[0].B8) != 0 && contains(blocks, "8") {
		//xml.AddrReg.Apartment = jsn[0].B8[0].RegAddrrecretFlatNum
		var addr = model_xml.AddrReg{}
		addr.Apartment = jsn[0].B8[0].RegAddrrecretFlatNum

		xml.AddrReg = &addr
	}

	if reflect.DeepEqual(xml, model_xml.BasePart{}) {
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
