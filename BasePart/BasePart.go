package BasePart

import (
	"go_pro/model_jsn"
	"go_pro/model_xml"
)

func BasePart(jsn model_jsn.Model_jsn, event string, blocks []string) *model_xml.BasePart {
	var xml = model_xml.BasePart{}

	if len(jsn[0].B8) != 0 && contains(blocks, "55") {
		xml.AddrReg.Apartment = jsn[0].B8[0].RegAddrrecretFlatNum
	}

	return &xml
}

func contains(a []string, v string) bool {
	for _, i := range a {
		if i == v {
			return true
		}
	}
	return false
}
