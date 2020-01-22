package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JsonResult(c *gin.Context, code int, err error, data interface{}) {
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  code,
			"error": err.Error(),
			"data":  data,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":  code,
			"error": nil,
			"data":  data,
		})
	}

}
