package core

import (
	"go-service/database"
	"go-service/initialize"
	"go-service/utils"
)

// RunningGinServe v1服务启动方法
func RunningGinServe() {
	config := utils.ReadYamlConfig()
	//	初始化数据库连接池
	database.InitDataBase()
	//	初始化路由中间件
	router := initialize.Routers()
	err := router.Run(":" + config.Project.Port)
	if err != nil {
		return
	}
}
