package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-service/services"
)

func UploadFile(ctx *gin.Context) {
	//	获取上传请求的头文件
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		fmt.Println("upload file error")
	}
	result, _ := services.CreateFileUpload(file, header)
	ctx.JSON(200, gin.H{
		"status": true,
		"msg":    "上传成功",
		"data":   result,
	})
}
