package initialize

import (
	"github.com/gin-gonic/gin"
	"go-service/middleware"
	"go-service/router"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	/* 启用跨域中间件 利用设置Request Header进行跨域配置 */
	Router.Use(middleware.Cors())
	apiGroup := Router.Group("v1")
	{
		router.InitUserRouter(apiGroup)
		router.InitRoleRouter(apiGroup)
		router.InitWsRouter(apiGroup)
		router.InitFileRouter(apiGroup)
	}
	return Router
}
