package utils

import "os"

// Exist function description : 判断文件或者目录是否存在
// params :
// return :
func IsExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
