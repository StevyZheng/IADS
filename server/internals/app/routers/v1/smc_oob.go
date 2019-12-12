package v1

import (
	"github.com/gin-gonic/gin"
	"iads/server/pkg/config"
	"iads/util"
)

type OobInfo struct {
	SN     string `json:"sn"`
	BmcMac string `json:"bmc_mac"`
}

func OobActiveCodeSN(c *gin.Context) {
	var oob OobInfo
	err := c.ShouldBindJSON(&oob)
	if oob.SN == "" {
		config.JsonRequest(c, -5, nil, nil)
		return
	}
	if err != nil {
		println(err.Error())
		config.JsonRequest(c, -3, nil, nil)
	} else {
		code, err := util.SmcOobActiveFunc(oob.BmcMac)
		if err != nil {
			println(err.Error())
			config.JsonRequest(c, -4, nil, nil)
		} else {
			config.JsonRequest(c, 1, code, err)
		}
	}
}

func OobActiveCode(c *gin.Context) {
	var oob OobInfo
	err := c.ShouldBindJSON(&oob)
	if err != nil {
		println(err.Error())
		config.JsonRequest(c, -3, nil, nil)
	} else {
		code, err := util.SmcOobActiveFunc(oob.BmcMac)
		if err != nil {
			println(err.Error())
			config.JsonRequest(c, -4, nil, nil)
		} else {
			config.JsonRequest(c, 1, code, err)
		}
	}
}
