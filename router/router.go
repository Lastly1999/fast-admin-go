package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-service/api/v1"
	"go-service/database"
)

func AppRouter() {
	database.InitDataBase()
	app := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true                                                                                                 //允许所有域名
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}                                                                      //允许请求的方法
	config.AllowHeaders = []string{"tus-resumable", "upload-length", "upload-metadata", "cache-control", "x-requested-with", "*"} //允许的Header
	app.Use(cors.New(config))
	v1Router := app.Group("v1")
	{
		// 用户模块
		userGroup := v1Router.Group("/user")
		{
			userGroup.POST("/login", v1.LoginAction)
		}
		// 权限模块
		roleGroup := v1Router.Group("/role")
		{
			roleGroup.GET("/menus", v1.FindtUserRoleMenus)
		}
		// websocket service
		v1Router.Group("/ws")
		{
			v1Router.GET("/v1ws", v1.WsPing)
		}
	}
	err := app.Run()
	if err != nil {
		return
	}
}
