/*****************************************************************************
 *  file name : SEngine.go
 *  author : Wu Yinghao
 *  email  : wyh817@gmail.com
 *
 *  file description : 引擎层，负责sql
 *
******************************************************************************/

package EngineLayer

import (
	dl "SparrowDB/DataLayer"
	"utils"
)

type SEngine struct {
	Pathname string `json:"dbpathname"`
	db       *dl.SDatabase
	Logger   *utils.Log4FE
}

// NewEngine function description : 构造引擎
// params :
// return :
func NewEngine(pathname string, logger *utils.Log4FE) *SEngine {

	this := &SEngine{Logger: logger, Pathname: pathname,dl:nil}

	return this

}

// OpenDB function description : 打开数据库
// params :
// return :
func (this *SEngine) OpenDB(dbname string) error {
    return nil
}


// CreateDB function description : 创建DB
// params : 
// return : 
func (this *SEngine) CreateDB(dbname string ) error {
    
    this.db = dl.NewSDatabase(dbname,this.Pathname,this.Logger)
    
    if thie.db == nil {
        return fmt.Errorf("Create DB Error ")
    } 
    
    return nil
    
}


func (this *SEngine) CreateTable(tablename string, fieldinfos []FieldMeta) error {
    
    return this.db.CreateTable(tablename,fieldinfos)
}



func (this *SEngine) AddData(tablename string ,content map[string]string) error {
    
    return this.db.AddData(tablename,content)
}




func parseSql(sql string)
