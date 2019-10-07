package file

import (
	"iads/lib/logging"
	"io/ioutil"
)

// 获取文件夹内各文件的文件名
func GetFolderSubFileName(path string) (fileNames []string, err error) {
	dirList, err := ioutil.ReadDir(path)
	if err != nil {
		logging.FatalPrintln("Read Dir " + path + " error.")
		return
	}
	for _, v := range dirList {
		fileNames = append(fileNames, v.Name())
	}
	return
}
