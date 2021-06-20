package router

import (
	"github.com/gin-gonic/gin"
	v1 "go-service/api/v1"
)

// InitUserRouter 用户路由集合
func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("/login", v1.LoginAction)
	}
}
