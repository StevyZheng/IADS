package routers

import (
	"github.com/gin-gonic/gin"
	"iads/server/manager/internals/app/routers/initer"
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
	apiRole := apiV1.Group("/role")
	initer.RoleRouterInit(apiRole)

	return router
}
