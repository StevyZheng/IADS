package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetCurrentDirectory() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return strings.Replace(dir, "\\", "/", -1)
}

func GetRootDir() string {
	file, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		file = fmt.Sprintf(".%s", string(os.PathSeparator))
	} else {
		file = fmt.Sprintf("%s%s", file, string(os.PathSeparator))
	}
	return file
}

func GetExecFilePath() string {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		file = fmt.Sprintf(".%s", string(os.PathSeparator))
	} else {
		file, _ = filepath.Abs(file)
	}
	return file
}

func ListFiles(path string) (files []string, err error) {
	files = make([]string, 0)
	fs, err := ioutil.ReadDir(path)
	for _, file := range fs {
		if !file.IsDir() {
			files = append(files, filepath.Join(path, file.Name()))
		}
	}
	return files, err
}
