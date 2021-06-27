package router

import (
	"github.com/gin-gonic/gin"
	v1 "go-service/handler/v1"
)

// InitCarRouter 汽车相关业务路由组
func InitCarRouter(Router *gin.RouterGroup) {
	carRouter := Router.Group("/car")
	{
		carRouter.POST("/list", v1.GetCarListById)
	}
}
