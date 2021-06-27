package router

import (
	"github.com/gin-gonic/gin"
	"go-service/database"
)

// InitUserRouter 用户路由集合
func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("/login", database.UserHandel.GetUserHandel)
	}
}
