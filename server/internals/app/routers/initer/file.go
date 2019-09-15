package initer

import (
	"github.com/gin-gonic/gin"
	"iads/server/internals/app/middleware"
	v1 "iads/server/internals/app/routers/v1"
)

func FileRouterInit(fileRouterGroup *gin.RouterGroup) {
	fileRouterGroup.Use(middleware.JWTAuth())
	{
		fileRouterGroup.POST("/upload", v1.FileUpload)
		fileRouterGroup.POST("/download", v1.FileDownload)
	}
}
