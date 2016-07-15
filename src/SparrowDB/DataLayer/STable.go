/*****************************************************************************
 *  file name : STable.go
 *  author : Wu Yinghao
 *  email  : wyh817@gmail.com
 *
 *  file description : 数据层
 *
******************************************************************************/

package DataLayer

import (
	"fmt"
	"utils"
)

const BT_TAIL string = ".bt"
const TB_DTL_TAIL string = ".tb.detail"

// FieldMeta 字段信息
type FieldMeta struct {
	Fieldname string `json:"fieldname"`
	FieldLen  int    `json:"fieldlen"`
	FieldType int    `json:"fieldtype"`
    Default   string `json:"default"`
	MkIdx     bool   `json:"makeindex"`
}

type STable struct {
	Tablename string               `json:"tablename"`
	Fields    map[string]FieldMeta `json:"fields"`
	FieldLen  int                  `json:"fieldlen"`
	Pathname  string
	btreeName string
	bt        *utils.BTreedb
	detail    *utils.Mmap
	Logger    *utils.Log4FE
}

// NewSTable function description : 新建数据库表
// params :
// return :
func NewSTable(tablename, pathname string, fieldsinfos []FieldMeta, logger *utils.Log4FE) *STable {

	this := &STable{Logger: logger, Pathname: pathname, Tablename: tablename,Fields:make(map[string]FieldMeta)}

	if utils.IsExist(pathname + "/" + tablename + TB_DTL_TAIL) {
		this.Logger.Error("[ERROR] STable[%v] is exist", tablename)
		return nil
	}
    
    for _,field := range fieldsinfos {
        if _,ok:=this.Fields[field.Fieldname];ok{
            this.Logger.Error("[ERROR] Field[%v] exist",field.Fieldname)
            return nil
        }
        this.Fields[field.Fieldname] = field
    }

	//创建表的索引，使用b+树索引
	if err := this.createIndex(); err != nil {
		this.Logger.Error("[ERROR] createIndex  %v", err)
		return nil
	}

	//创建detail文件
	if err := this.createDetail(); err != nil {
		this.Logger.Error("[ERROR] createDetail  %v", err)
		return nil
	}

	this.Logger.Info("[INFO] STable[%v] Create ok", tablename)

	return this
}

// createIndex function description : 创建索引
// params :
// return :
func (this *STable) createIndex() error {

	//初始化b+树索引
	this.btreeName = this.Pathname + "/" + this.Tablename + BT_TAIL
	this.bt = utils.NewBTDB(this.btreeName)

	if this.bt == nil {
		this.Logger.Error("[ERROR] make b+tree error %v", this.btreeName)
		return fmt.Errorf("[ERROR] make b+tree error %v", this.btreeName)
	}

	for k, v := range this.Fields {
        
        this.FieldLen = this.FieldLen + v.FieldLen
        

		if v.MkIdx {
			this.bt.AddBTree(k)
		}

	}

	return nil

}


// createDetail function description : 创建表的详情文件
// params : 
// return : 
func (this *STable) createDetail() error {

	var err error

	this.detail, err = utils.NewMmap(this.Pathname+"/"+this.Tablename+TB_DTL_TAIL, utils.MODE_CREATE)

	if err != nil {
		this.Logger.Error("[ERROR] make detail error %v", err)
		return err
	}

	return nil

}

func (this *STable) AddData(content map[string]string) error {
    return nil
}

func (this *STable) DeleteData(docid utils.DocID) error {

	return nil
}

func (this *STable) UpdateData(docid utils.DocID, content map[string]string) error {

	return nil
}

func (this *STable) FindData(field, value string) []utils.DocID {
    return nil
}

func (this *STable) FindDocId(docid utils.DocID) map[string]string {
    return nil
}
