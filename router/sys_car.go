package router

import (
	"github.com/gin-gonic/gin"
	v1 "go-service/api/v1"
)

// InitCarRouter 汽车相关业务路由组
func InitCarRouter(Router *gin.RouterGroup) {
	carRouter := Router.Group("/car")
	{
		carRouter.GET("/list/:typeId", v1.GetCarListById)
	}
}
