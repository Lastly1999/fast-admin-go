package router

import (
	"github.com/gin-gonic/gin"
	v1 "go-service/api/v1"
)

func InitFileRouter(Router *gin.RouterGroup) {
	FileRouter := Router.Group("upload")
	{
		FileRouter.POST("/file", v1.UploadFile)
	}
}
