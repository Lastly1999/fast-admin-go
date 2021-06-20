package initialize

import (
	"github.com/gin-gonic/gin"
	"go-service/middleware"
	"go-service/router"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.Use(middleware.Cors())
	apiGroup := Router.Group("v1")
	{
		router.InitUserRouter(apiGroup)
		router.InitRoleRouter(apiGroup)
	}
	return Router
}
