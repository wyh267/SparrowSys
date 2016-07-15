package DataLayer

import (
	"encoding/json"
	"fmt"
	"testing"
	"utils"
)

const tbinfo string = `[
    {
        "fieldname":"id",
        "fieldlen":10,
        "fieldtype":0,
        "makeindex":true
    },
    {
        "fieldname":"name",
        "fieldlen":10,
        "fieldtype":0,
        "makeindex":true
    },
    {
        "fieldname":"age",
        "fieldlen":4,
        "fieldtype":0,
        "makeindex":false
    }
    
    
]`

func TestNewDB(t *testing.T) {

	logger, _ := utils.New("test_db")

	db := NewSDatabase("testdb", "./testdata", logger)

	var feilds []FieldMeta
	err := json.Unmarshal([]byte(tbinfo), feilds)
	if err != nil {
		fmt.Printf("error\n")
	}


    fields := make([]FieldMeta,0) 
    fields = append(fields,FieldMeta{FieldLen:10,Fieldname:"id",FieldType:0,MkIdx:true})
    fields = append(fields,FieldMeta{FieldLen:10,Fieldname:"name",FieldType:0,MkIdx:true})
    fields = append(fields,FieldMeta{FieldLen:10,Fieldname:"age",FieldType:0,MkIdx:false})



	err = db.CreateTable("biao", fields)
	if err != nil {
		fmt.Printf("biao err\n")
	}

}
