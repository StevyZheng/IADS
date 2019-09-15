package initer

import (
	"github.com/gin-gonic/gin"
	"iads/server/internals/app/middleware"
	v1 "iads/server/internals/app/routers/v1"
)

func RoleRouterInit(roleRouterGroup *gin.RouterGroup) {
	roleRouterGroup.Use(middleware.JWTAuth())
	{
		roleRouterGroup.GET("/list", v1.RoleList)
		roleRouterGroup.GET("/get/:role_name", v1.RoleGetFromName)
		roleRouterGroup.POST("/add", v1.RoleCreate)
		roleRouterGroup.POST("/del/:role_name", v1.RoleDestroyFromRoleName)
	}
}
