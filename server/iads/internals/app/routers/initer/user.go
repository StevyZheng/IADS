package initer

import (
	"github.com/gin-gonic/gin"
	middleware2 "iads/server/iads/internals/app/middleware"
	v12 "iads/server/iads/internals/app/routers/v1"
)

func UserRouterInit(userRouterGroup *gin.RouterGroup) {
	userRouterGroup.Use(middleware2.JWTAuth())
	{
		userRouterGroup.GET("/list", v12.UserList)
		userRouterGroup.GET("/get/:username", v12.UserGetFromName)
		userRouterGroup.POST("/add", v12.UserCreate)
		userRouterGroup.POST("/del", v12.UserDestroyFromUserName)
		userRouterGroup.POST("/del/:username", v12.UserDestroy)
	}
}
