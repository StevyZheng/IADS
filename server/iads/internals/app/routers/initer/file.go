package initer

import (
	"github.com/gin-gonic/gin"
	middleware2 "iads/server/iads/internals/app/middleware"
	v12 "iads/server/iads/internals/app/routers/v1"
)

func FileRouterInit(fileRouterGroup *gin.RouterGroup) {
	fileRouterGroup.Use(middleware2.JWTAuth())
	{
		fileRouterGroup.POST("/upload", v12.FileUpload)
		fileRouterGroup.POST("/download", v12.FileDownload)
	}
}
