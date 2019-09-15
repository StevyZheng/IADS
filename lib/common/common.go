package common

import (
	"bytes"
	"github.com/mholt/archiver"
	"os/exec"
)

func ExecShellLinux(cmd string) (string, error) {
	ret := exec.Command("/bin/bash", "-c", cmd)
	var out bytes.Buffer
	ret.Stdout = &out
	err := ret.Run()
	return out.String(), err
}

func ExecShellWin(cmd string) (string, error) {
	ret := exec.Command("cmd", "/C", cmd)
	var out bytes.Buffer
	ret.Stdout = &out
	err := ret.Run()
	return out.String(), err
}

func ZipFiles(fileArr []string, tgzName string) error {
	zip := archiver.Zip{
		CompressionLevel:       9,
		MkdirAll:               true,
		SelectiveCompression:   false,
		ContinueOnError:        false,
		OverwriteExisting:      false,
		ImplicitTopLevelFolder: false,
	}
	err := zip.Archive(fileArr, tgzName)
	return err
}

func ZipPath(path string, zipName string) error {
	err := ZipFiles([]string{path}, zipName)
	return err
}

func UnZip(zipName string, pathName string) error {
	err := archiver.Unarchive(zipName, pathName)
	return err
}

func CheckErr(err error) {
	if err != nil {
		println(err)
	}
}
