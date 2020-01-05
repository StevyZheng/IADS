package initer

import (
	"github.com/gin-gonic/gin"
	v1 "iads/server/manager/internals/app/routers/v1"
)

func RoleRouterInit(roleRouterGroup *gin.RouterGroup) {
	/*roleRouterGroup.Use(middleware.JWTAuth())
	{
		roleRouterGroup.GET("/list", v1.RoleList)
	}*/
	roleRouterGroup.GET("/list", v1.RoleList)
}
