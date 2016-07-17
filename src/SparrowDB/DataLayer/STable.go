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
	"encoding/binary"
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
	Tablename  string               `json:"tablename"`
	Fields     map[string]FieldMeta `json:"fields"`
	FieldInfos []FieldMeta          `json:"fieldinfos"`
	RecordLen  int                  `json:"RecordLen"`
	MaxCount   int64                `json:"maxcount"`
	Pathname   string               `json:"pathname"`
	btreeName  string
	bt         *utils.BTreedb
	detail     *utils.Mmap
	Logger     *utils.Log4FE
}

// NewSTable function description : 新建数据库表
// params :
// return :
func NewSTable(tablename, pathname string, fieldsinfos []FieldMeta, logger *utils.Log4FE) *STable {

	this := &STable{MaxCount: 0, Logger: logger, Pathname: pathname, Tablename: tablename, FieldInfos: fieldsinfos, Fields: make(map[string]FieldMeta)}

	if utils.IsExist(pathname + "/" + tablename + TB_DTL_TAIL) {
		this.Logger.Error("[ERROR] STable[%v] is exist", tablename)
		return nil
	}

	for _, field := range fieldsinfos {
		if _, ok := this.Fields[field.Fieldname]; ok {
			this.Logger.Error("[ERROR] Field[%v] exist", field.Fieldname)
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

		this.RecordLen = this.RecordLen + v.FieldLen + 4

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

// AddData function description : 添加数据
// params :
// return :
func (this *STable) AddData(content map[string]string) error {

	inbytes := make([]byte, this.RecordLen)
	point := uint32(0)
	var value string

	for _, fvalue := range this.FieldInfos {

		if _, ok := content[fvalue.Fieldname]; !ok {

			value = fvalue.Default

		} else {
			value = content[fvalue.Fieldname]
		}

		lens := uint32(len(value))
		binary.LittleEndian.PutUint32(inbytes[point:point+4], lens)
		point += 4

		dst := inbytes[point : point+lens]
		copy(dst, []byte(value))
		point += uint32(fvalue.FieldLen)
        
        
        //如果有索引要求
        if fvalue.MkIdx {
            //this.bt.A
            
            
            
        }
        
        

	}
	this.detail.AppendRecord(inbytes)
	this.MaxCount++
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

func (this *STable) FindDocId(docid int64) map[string]string {

	if docid >= this.MaxCount {
		return nil
	}

	res := make(map[string]string)

	offset := docid * int64(this.RecordLen)

	outbytes := this.detail.ReadRecord(offset, uint32(this.RecordLen))
	point := uint32(0)

	for _, fvalue := range this.FieldInfos {

		reallen := binary.LittleEndian.Uint32(outbytes[point : point+4])
		point += 4
		value := string(outbytes[point : point+reallen])
		point += uint32(fvalue.FieldLen)
        res[fvalue.Fieldname]=value

	}

	return res
}
