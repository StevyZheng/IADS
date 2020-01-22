package initer

import (
	"github.com/gin-gonic/gin"
	"iads/server/manager/internals/app/middleware"
	v1 "iads/server/manager/internals/app/routers/v1"
)

func UserRouterInit(roleRouterGroup *gin.RouterGroup) {
	roleRouterGroup.Use(middleware.JWTAuth())
	{
		roleRouterGroup.GET("/list", v1.UserList)
		roleRouterGroup.POST("/add_one", v1.UserAddOne)
		roleRouterGroup.POST("/del_one", v1.UserDeleteFromName)
		roleRouterGroup.POST("/update_one", v1.UserUpdateOneFromName)
	}
}
