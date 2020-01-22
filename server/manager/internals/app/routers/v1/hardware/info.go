package hardware

import (
	"github.com/gin-gonic/gin"
	"iads/lib/linux/hardware"
	v1 "iads/server/manager/internals/app/routers/v1"
)

func CpuInfo(c *gin.Context) {
	cpuInfo := new(hardware.CpuHwInfo)
	if err := cpuInfo.GetCpuHwInfo(); err != nil {
		v1.JsonResult(c, 409, err, nil)
	}
	v1.JsonResult(c, 200, nil, cpuInfo)
}

func MbInfo(c *gin.Context) {
	mbInfo := new(hardware.MotherboradInfo)
	mbInfo.GetMbInfo()
	v1.JsonResult(c, 200, nil, mbInfo)
}

func DiskInfo(c *gin.Context) {
	disks, err := hardware.Disk{}.DiskList()
	if err != nil {
		v1.JsonResult(c, 209, nil, err)
	} else {
		v1.JsonResult(c, 200, nil, disks)
	}
}
