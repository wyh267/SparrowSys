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
	"fmt"
	"os"
	"utils"
)

type SDatabase struct {
	Name     string `json:"dbname"`
	Pathname string `json:"pathname"`
	Fullname string `json:"fullname`
	tables   map[string]*STable

	Logger *utils.Log4FE
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

// initDatabase function description : 初始化数据库
// params :
// return :
func (this *SDatabase) initDatabase() error {

	//初始化table
	this.tables = make(map[string]*STable)

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

	table := NewSTable(tablename, this.Fullname, fieldinfos,this.Logger)

	if table == nil {
		this.Logger.Error("[ERROR] Create Table[%v] Fail", tablename)
		return fmt.Errorf("[ERROR] Create Table[%v] Fail", tablename)
	}

	this.tables[tablename] = table

	return nil

}

func (this *SDatabase) AddData(content map[string]string) error {

	return nil

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

func (this *SDatabase) FindDocId(docid utils.DocID) map[string]string {

	return nil
}
