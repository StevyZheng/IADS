package initer

import (
	"github.com/gin-gonic/gin"
	middleware2 "iads/server/iads/internals/app/middleware"
	v12 "iads/server/iads/internals/app/routers/v1"
)

func OobRouterInit(OobRouterGroup *gin.RouterGroup) {
	OobRouterGroup.Use(middleware2.JWTAuth())
	{
		OobRouterGroup.POST("/active", v12.OobActiveCode)
	}
}
