package core

import (
	"go-service/database"
	"go-service/initialize"
)

// RunningGinServe v1服务启动方法
func RunningGinServe() {
	//	config := utils.ReadYamlConfig()
	//	初始化数据库连接池
	database.InitDataBase()
	//	初始化路由中间件
	router := initialize.Routers()
	//	serve := ":" + config.Project.Port
	//	":"+config.Project.Port
	//	不需要模板渲染
	//	router.LoadHTMLGlob("templates/*")
	err := router.Run()
	if err != nil {
		return
	}
}
