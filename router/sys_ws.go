package router

import (
	"github.com/gin-gonic/gin"
	v1 "go-service/handler/v1"
)

// InitWsRouter 用户路由集合
func InitWsRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("ws")
	{
		UserRouter.GET("/v1ws", v1.WsPing)
	}
}
