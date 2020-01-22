package initer

import (
	"github.com/gin-gonic/gin"
	"iads/server/manager/internals/app/middleware"
	"iads/server/manager/internals/app/routers/v1/hardware"
)

func HardwareRouterInit(hardwareRouterGroup *gin.RouterGroup) {
	hardwareRouterGroup.Use(middleware.JWTAuth())
	{
		hardwareRouterGroup.GET("/cpuinfo", hardware.CpuInfo)
		hardwareRouterGroup.GET("/mbinfo", hardware.MbInfo)
		hardwareRouterGroup.GET("/disk_list", hardware.DiskInfo)
	}
}
