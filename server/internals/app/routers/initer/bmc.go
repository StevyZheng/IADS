package initer

import (
	"github.com/gin-gonic/gin"
	"iads/server/internals/app/middleware"
	v1 "iads/server/internals/app/routers/v1"
)

func OobRouterInit(OobRouterGroup *gin.RouterGroup) {
	OobRouterGroup.Use(middleware.JWTAuth())
	{
		OobRouterGroup.POST("/active", v1.OobActiveCode)
	}
}
