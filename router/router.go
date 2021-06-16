package router

import (
	"github.com/gin-gonic/gin"
	"go-service/api/v1"
	"go-service/database"
)

func AppRouter()  {
	app := gin.Default()
	database.InitDataBase()
	v1Router := app.Group("v1")
	{
		v1Router.GET("/",v1.Test)
		v1Router.POST("/body",v1.PostBody)
	}
	err := app.Run()
	if err != nil {
		return 
	}
}
