package router

import (
	"github.com/gin-gonic/gin"
	"go-service/api/v1"
	"go-service/database"
)

func AppRouter() {
	database.InitDataBase()
	app := gin.Default()
	v1Router := app.Group("v1")
	{
		v1Router.POST("/login", v1.LoginAction)
		v1Router.POST("/body", v1.PostBody)
		v1Router.GET("/v1ws", v1.WsPing)
		// 为ws服务开启一个新线程
		//go v1Router.GET("/v1ws", v1.WsPing)
	}
	err := app.Run()
	if err != nil {
		return
	}
}
