package routers

import (
	"github.com/gin-gonic/gin"
	initer2 "iads/server/iads/internals/app/routers/initer"
	v12 "iads/server/iads/internals/app/routers/v1"
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
	apiV1.POST("/login", v12.Login)
	apiV1.POST("/register", v12.UserCreate)

	//apiDoc := apiV1.Group("/doc")
	//routers.ApiRouterInit(apiDoc)

	apiRole := apiV1.Group("/role")
	initer2.RoleRouterInit(apiRole)

	apiUser := apiV1.Group("/user")
	initer2.UserRouterInit(apiUser)

	apiFile := apiV1.Group("/file")
	initer2.FileRouterInit(apiFile)

	apiBmc := apiV1.Group("/bmc")
	initer2.OobRouterInit(apiBmc)

	return router
}
