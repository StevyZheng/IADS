package routers

import (
	"github.com/gin-gonic/gin"
	"iads/server/internals/app/routers/initer"
	v1 "iads/server/internals/app/routers/v1"
	"net/http"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	//router.NoRoute(api.NotFound)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"app":     "rs",
			"version": "1.0",
		})
	})

	apiV1 := router.Group("/api/v1")
	apiV1.POST("/login", v1.Login)
	apiV1.POST("/register", v1.UserCreate)

	//apiDoc := apiV1.Group("/doc")
	//routers.ApiRouterInit(apiDoc)

	apiRole := apiV1.Group("/role")
	initer.RoleRouterInit(apiRole)

	apiUser := apiV1.Group("/user")
	initer.UserRouterInit(apiUser)

	//apiFile := apiV1.Group("file")
	//routers.FileRouterInit(apiFile)

	apiFile := apiV1.Group("file")
	initer.FileRouterInit(apiFile)

	return router
}
