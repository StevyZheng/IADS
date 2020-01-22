package routers

import (
	"github.com/gin-gonic/gin"
	"iads/server/manager/internals/app/routers/initer"
	v1 "iads/server/manager/internals/app/routers/v1"
	"net/http"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	//router.NoRoute(api.NotFound)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"app":     "manager",
			"version": "1.0",
		})
	})

	apiV1 := router.Group("/api/v1")
	apiV1.POST("/login", v1.Login)

	apiRole := apiV1.Group("/role")
	initer.RoleRouterInit(apiRole)
	apiUser := apiV1.Group("/user")
	initer.UserRouterInit(apiUser)
	apiHardware := apiV1.Group("/hardware")
	initer.HardwareRouterInit(apiHardware)

	return router
}
