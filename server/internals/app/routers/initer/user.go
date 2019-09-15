package initer

import (
	"github.com/gin-gonic/gin"
	"iads/server/internals/app/middleware"
	v1 "iads/server/internals/app/routers/v1"
)

func UserRouterInit(userRouterGroup *gin.RouterGroup) {
	userRouterGroup.Use(middleware.JWTAuth())
	{
		userRouterGroup.GET("/list", v1.UserList)
		userRouterGroup.GET("/get/:user_name", v1.UserGetFromName)
		userRouterGroup.POST("/add", v1.UserCreate)
		userRouterGroup.POST("/del", v1.UserDestroyFromUserName)
		userRouterGroup.POST("/del/:user_name", v1.UserDestroy)
	}
}
