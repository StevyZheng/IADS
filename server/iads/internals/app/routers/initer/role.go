package initer

import (
	"github.com/gin-gonic/gin"
	middleware2 "iads/server/iads/internals/app/middleware"
	v12 "iads/server/iads/internals/app/routers/v1"
)

func RoleRouterInit(roleRouterGroup *gin.RouterGroup) {
	roleRouterGroup.Use(middleware2.JWTAuth())
	{
		roleRouterGroup.GET("/list", v12.RoleList)
		roleRouterGroup.GET("/get/:role_name", v12.RoleGetFromName)
		roleRouterGroup.POST("/add", v12.RoleCreate)
		roleRouterGroup.POST("/del/:role_name", v12.RoleDestroyFromRoleName)
	}
}
