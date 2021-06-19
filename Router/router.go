package Router

import (
	"github.com/gin-gonic/gin"
	"go-service/Api/v1"
	"go-service/DataBase"
)

func AppRouter() {
	DataBase.InitDataBase()
	app := gin.Default()
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
