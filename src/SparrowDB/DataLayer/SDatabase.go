/*****************************************************************************
 *  file name : SDatabase.go
 *  author : Wu Yinghao
 *  email  : wyh817@gmail.com
 *
 *  file description : 数据层
 *
******************************************************************************/

package DataLayer

import (
	"encoding/json"
	"fmt"
	"os"
	"utils"
)

type TableInfo struct {
	Tablename string `json:"tablename"`
}

type SDatabase struct {
	Name       string             `json:"dbname"`
	Pathname   string             `json:"pathname"`
	Fullname   string             `json:"fullname`
	TableNames []string           `json:"tables"`
	tables     map[string]*STable `json:"-"`

	Logger *utils.Log4FE `json:"-"`
}

//NewSTable(tablename, pathname string, fieldsinfo map[string]FieldMeta, logger *utils.Log4FE) *STable
// NewSDatabase function description : 新建数据库
// params :
// return :
func NewSDatabase(dbname, dbpath string, logger *utils.Log4FE) *SDatabase {

	this := &SDatabase{Logger: logger, Pathname: dbpath, Name: dbname}

	this.Fullname = dbpath + "/" + dbname
	if utils.IsExist(this.Fullname) {
		this.Logger.Error("[ERROR] SDatabase[%v] is exist", dbname)
		return nil
	}

	os.MkdirAll(this.Fullname, 0777)

	if err := this.initDatabase(); err != nil {
		return this
	}

	this.Logger.Info("[INFO] SDatabase[%v] Create ok", dbname)

	return this
}

func OpenSDatabase(dbname, dbpath string, logger *utils.Log4FE) *SDatabase {

	this := &SDatabase{Logger: logger, Pathname: dbpath, Name: dbname}

	this.Fullname = dbpath + "/" + dbname


	buffer, err := utils.ReadFromJson(this.Fullname + "/_dbinfo.meta")
	if err != nil {
		return nil
	}

	err = json.Unmarshal(buffer, this)
	if err != nil {
		return nil
	}
    
    for _,tablename := range this.TableNames {
        
        //TODO
        
        
    }
    
    return this
    

}

// initDatabase function description : 初始化数据库
// params :
// return :
func (this *SDatabase) initDatabase() error {

	//初始化table
	this.tables = make(map[string]*STable)
	this.TableNames = make([]string, 0)

	return nil

}

// CreateTable function description : 创建数据库表
// params :
// return :
func (this *SDatabase) CreateTable(tablename string, fieldinfos []FieldMeta) error {

	if _, ok := this.tables[tablename]; ok {
		this.Logger.Error("[ERROR] Table[%v] Create Error", tablename)
		return fmt.Errorf("[ERROR] Table[%v] Create Error", tablename)

	}

	table := NewSTable(tablename, this.Fullname, fieldinfos, this.Logger)

	if table == nil {
		this.Logger.Error("[ERROR] Create Table[%v] Fail", tablename)
		return fmt.Errorf("[ERROR] Create Table[%v] Fail", tablename)
	}

	this.tables[tablename] = table
	this.TableNames = append(this.TableNames, tablename)

	utils.WriteToJson(this, this.Fullname+"/_dbinfo.meta")

	return nil

}

func (this *SDatabase) AddData(tablename string, content map[string]string) error {

	if _, ok := this.tables[tablename]; ok {

		return this.tables[tablename].AddData(content)

	}

	return fmt.Errorf("[ERROR] Table[%v] Not Found", tablename)

}

func (this *SDatabase) DeleteData(docid utils.DocID) error {

	return nil
}

func (this *SDatabase) UpdateData(docid utils.DocID, content map[string]string) error {

	return nil
}

func (this *SDatabase) FindData(field, value string) []utils.DocID {

	return nil

}

func (this *SDatabase) FindDocId(tablename string, docid int64) map[string]string {

	if _, ok := this.tables[tablename]; ok {

		return this.tables[tablename].FindDocId(docid)

	}

	this.Logger.Error("[ERROR] Table[%v] Not Found", tablename)

	return nil
}
