package routes

import (
	"github.com/gin-gonic/gin"
	"webtry/controller"
)

/**
 * @Author: lbh
 * @Date: 2021/4/9
 * @Description:
 */
func CollectRoute(r *gin.Engine) *gin.Engine{
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)

	return r
}