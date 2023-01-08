package main

import (
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"go_pro/InformationPart"
	"go_pro/TitlePart"
	"go_pro/model_jsn"
	"go_pro/model_xml"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "marcus123"
	dbname   = "event_db"
)

var fch = model_xml.Fch{}

func main() {
	//////////////////////////// connect to database ///////////////////////////////
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	rows, err := db.Query(`select * from data2send.data2send`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	//////////////////////////////// rows handle: ///////////////////////////////////
	for rows.Next() {
		var event int
		var row_jsn json.Number
		err = rows.Scan(&event, &row_jsn)
		if err != nil {
			panic(err)
		}
		JsnToXml(row_jsn)
	}
	//////////////////////////////// export to file: ////////////////////////////////
	out, _ := xml.MarshalIndent(fch, " ", "  ")
	err = os.WriteFile("output.xml", []byte(xml.Header+string(out)), 0644)
	if err != nil {
		panic(err)
	}
	//////////////////////////////////////////////////////////////////////////////////
}

func JsnToXml(row_jsn json.Number) {
	jsn := model_jsn.Model_jsn{}
	err := json.Unmarshal([]byte(row_jsn), &jsn)
	if err != nil {
		panic(err)
	}
	fch.Info = append(fch.Info, model_xml.Info{})
	j := len(fch.Info) - 1
	fch.Info[j].Action = jsn[0].ActionType
	fch.Info[j].Event = jsn[0].EvtCode
	fch.Info[j].Recnumber = strconv.Itoa(jsn[0].Rn)

	var blocks []string
	switch jsn[0].EvtCode {
	case "11":
		blocks = append(blocks, "55")
	case "13":
		blocks = append(blocks, "55", "57")
	}

	fch.Info[j].TitlePart = TitlePart.TitlePart(jsn)
	fch.Info[j].InformationPart = InformationPart.InformationPart(jsn, jsn[0].EvtCode, blocks)

}

func Contains(a []string, v string) bool {
	for _, i := range a {
		if i == v {
			return true
		}
	}
	return false
}
