package file

import (
	"bufio"
	"iads/lib/logging"
	"io"
	"io/ioutil"
	"os"
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

func IfFileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func ReadFileAsLine(path string) (lines []string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err == io.EOF {
				err = nil
				break
			}
		} else {
			lines = append(lines, string(line))
		}
	}
	return lines, err
}
