package util

import (
	"iads/lib/linux/hardware"
	"os"
	"path/filepath"
)

func RebootFunc() (err error) {
	fileName := "/etc/rc.d/rc.local"
	dir, _ := os.Executable()
	currentDir := filepath.Dir(dir)
	me := filepath.Join(currentDir, "iads")
	err = os.Chmod(fileName, 0777)
	fd, _ := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND, 0644)
	buf := "\n" + me + " reboot &"
	_, err = fd.WriteString(buf)
	fd.Close()
	return err
}

func SmcOobActiveFunc(bmcMac string) (code string, err error) {
	return hardware.OutActivationCode(bmcMac)
}
