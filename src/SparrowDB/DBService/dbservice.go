/*****************************************************************************
 *  file name : dbservice.go
 *  author : Wu Yinghao
 *  email  : wyh817@gmail.com
 *
 *  file description :
 *
******************************************************************************/

//package DBService
package main

import (
	"fmt"
	"utils"
	dl "SparrowDB/DataLayer"
)

func main() {

	fmt.Printf("Start DB ....")
	logger,_ := utils.New("logname")
	
	logger.Info("info...")
	db:=dl.NewSDatabase("testdb","./",logger)
	db.AddData(nil)
	
}
