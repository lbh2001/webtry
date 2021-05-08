package filter

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webtry/utils"
)

/**
 * @Author: lbh
 * @Date: 2021/5/8
 * @Description:
 */

func TokenFilterMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authority")
		if tokenString == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "No Token!",
			})
			c.Abort()
			return
		}
		if token, err := utils.ParseToken(tokenString); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "error:" + err.Error(),
			})
			c.Abort()
			return
		} else {
			c.Set("phone", token.Phone)
			c.Next()
		}

	}
}
