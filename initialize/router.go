package initialize

import (
	"github.com/gin-gonic/gin"
	"go-service/middleware"
	"go-service/router"
	"net/http"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	//	配置静态文件路径
	Router.StaticFS("/static", http.Dir("./static"))
	//	启用跨域中间件 利用设置Request Header进行跨域配置
	Router.Use(middleware.Cors())
	apiGroup := Router.Group("v1")
	{
		router.InitUserRouter(apiGroup)
		router.InitRoleRouter(apiGroup)
		router.InitWsRouter(apiGroup)
		router.InitFileRouter(apiGroup)
		router.InitCarRouter(apiGroup)
	}
	return Router
}
