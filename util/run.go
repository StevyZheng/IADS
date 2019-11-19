package util

import (
	"iads/server/internals/app/routers"
	"iads/server/internals/pkg/models/sys"
	"os"
	"path/filepath"
)

func ServerRunFunc() {
	println("iads api server is running...")
	//defer database.DBE.Close()
	sys.ModelInit()
	router := routers.InitRouter()
	_ = router.Run(":80")
}

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
