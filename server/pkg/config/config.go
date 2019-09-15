package config

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Conf struct {
	DBType       string
	UploadPath   string
	DownloadPath string
	OSType       string
}

var CodeMap map[int64]string
var ConfValue Conf

func init() {
	CodeMap = map[int64]string{
		-1: "操作失败",
		-2: "未找到相关信息",
		-3: "json解析失败",
		1:  "操作成功",
	}
	ConfValue.DBType = "pgsql"
	ConfValue.OSType = "windows"
	ConfValue.UploadPath = "/home/user/upload"
	ConfValue.DownloadPath = "/home/user/upload/"
}

func JsonRequest(c *gin.Context, code int64, data interface{}, err error) {
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": CodeMap[code],
		"error":   err,
		"data":    data,
	})
}

func StrToUint(strNumber string, value interface{}) (err error) {
	var number interface{}
	number, err = strconv.ParseUint(strNumber, 10, 64)
	switch v := number.(type) {
	case uint64:
		switch d := value.(type) {
		case *uint64:
			*d = v
		case *uint:
			*d = uint(v)
		case *uint16:
			*d = uint16(v)
		case *uint32:
			*d = uint32(v)
		case *uint8:
			*d = uint8(v)
		}
	}
	return
}
